"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { toast } from "sonner";
import dayjs from "dayjs";
import { IconPlus, IconTrash } from "@tabler/icons-react";
import { airportService, planeService, parameterService, flightService } from "@/services/api";
import { Airport, Plane, Parameter, StopAirport, FlightFormData } from "@/types/flight";

export default function FlightPage() {
  const router = useRouter();
  const [airports, setAirports] = useState<Airport[]>([]);
  const [planes, setPlanes] = useState<Plane[]>([]);
  const [stopAirports, setStopAirports] = useState<StopAirport[]>([]);
  const [parameter, setParameter] = useState<Parameter>({
    min_flight_duration: 0,
    max_intermediate_stops: 0,
    min_intermediate_stop_duration: 0,
    max_intermediate_stop_duration: 0,
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [airportsData, planesData, parameterData] = await Promise.all([
          airportService.getAirports(),
          planeService.getPlanes(),
          parameterService.getParameters(),
        ]);
        setAirports(airportsData);
        setPlanes(planesData);
        setParameter(parameterData);
      } catch (error) {
        toast.error("Không thể tải dữ liệu. Vui lòng thử lại sau.");
      }
    };
    fetchData();
  }, []);

  const onFinish = async (values: FlightFormData) => {
    const {
      departureAirport,
      arrivalAirport,
      duration,
      departureTime,
      flightCode,
      price,
    } = values;

    if (departureAirport === arrivalAirport) {
      toast.error("Sân bay đi và đến không được trùng nhau");
      return;
    }

    const middleAirports = stopAirports.map((a) => a.airport).filter(Boolean);

    const hasInvalidMiddle = middleAirports.some(
      (a) => a === departureAirport || a === arrivalAirport
    );

    if (hasInvalidMiddle) {
      toast.error("Sân bay trung gian không được trùng sân bay đi hoặc đến");
      return;
    }

    const hasDuplicateMiddle = new Set(middleAirports).size !== middleAirports.length;
    if (hasDuplicateMiddle) {
      toast.error("Không được chọn 2 sân bay trung gian giống nhau");
      return;
    }

    const hasInvalidStopTime = stopAirports.some(
      (stop) =>
        stop.time < parameter.min_intermediate_stop_duration ||
        stop.time > parameter.max_intermediate_stop_duration
    );
    if (hasInvalidStopTime) {
      toast.error("Thời gian dừng phải từ 10 đến 20 phút");
      return;
    }

    const formattedDepartureTime = departureTime.format("YYYY-MM-DD HH:mm:ss");

    const flightData = {
      arrival_airport: arrivalAirport,
      base_price: price,
      departure_airport: departureAirport,
      departure_date: formattedDepartureTime,
      duration: duration,
      intermediate_stops: stopAirports.map((stop, index) => ({
        note: stop.note,
        stop_airport: stop.airport,
        stop_duration: stop.time,
        stop_order: index + 1,
      })),
      plane_code: flightCode,
    };

    try {
      await flightService.addFlight(flightData);
      toast.success("Chuyến bay đã được thêm thành công");
      router.push("/dashboard");
    } catch (error: any) {
      toast.error(error.message || "Có lỗi xảy ra khi thêm chuyến bay");
    }
  };

  const addStopAirport = () => {
    if (stopAirports.length >= parameter.max_intermediate_stops) return;
    setStopAirports([
      ...stopAirports,
      { key: Date.now().toString(), airport: "", time: 0, note: "" },
    ]);
  };

  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "calc(var(--spacing) * 72)",
          "--header-height": "calc(var(--spacing) * 12)",
        } as React.CSSProperties
      }
    >
      <AppSidebar variant="inset" />
      <SidebarInset>
        <SiteHeader />
        <div className="flex flex-1 flex-col">
          <div className="@container/main flex flex-1 flex-col gap-2">
            <div className="flex flex-col gap-4 py-4 md:gap-6 md:py-6">
              <div className="px-4 lg:px-6">
                <Card>
                  <CardHeader>
                    <CardTitle className="text-2xl text-center">THÊM LỊCH CHUYẾN BAY</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <form onSubmit={(e) => {
                      e.preventDefault();
                      const formData = new FormData(e.currentTarget);
                      const departureTimeValue = formData.get('departureTime');
                      if (!departureTimeValue || typeof departureTimeValue !== 'string') {
                        toast.error("Vui lòng chọn thời gian khởi hành");
                        return;
                      }
                      onFinish({
                        departureAirport: formData.get('departureAirport') as string,
                        arrivalAirport: formData.get('arrivalAirport') as string,
                        duration: Number(formData.get('duration')),
                        departureTime: dayjs(departureTimeValue),
                        flightCode: formData.get('flightCode') as string,
                        price: Number(formData.get('price')),
                      });
                    }} className="space-y-4">
                      <div className="space-y-2">
                        <Label htmlFor="departureAirport">Sân bay đi</Label>
                        <Select name="departureAirport" required>
                          <SelectTrigger>
                            <SelectValue placeholder="Chọn sân bay" />
                          </SelectTrigger>
                          <SelectContent>
                            {airports.map((airport) => (
                              <SelectItem key={airport.airport_code} value={airport.airport_code}>
                                {airport.airport_name}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </div>

                      <div className="space-y-2">
                        <Label htmlFor="arrivalAirport">Sân bay đến</Label>
                        <Select name="arrivalAirport" required>
                          <SelectTrigger>
                            <SelectValue placeholder="Chọn sân bay" />
                          </SelectTrigger>
                          <SelectContent>
                            {airports.map((airport) => (
                              <SelectItem key={airport.airport_code} value={airport.airport_code}>
                                {airport.airport_name}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </div>

                      <div className="space-y-2">
                        <Label htmlFor="duration">Thời gian bay (phút)</Label>
                        <Input
                          type="number"
                          name="duration"
                          min={parameter.min_flight_duration}
                          required
                        />
                      </div>

                      <div className="space-y-2">
                        <Label htmlFor="departureTime">Ngày khởi hành</Label>
                        <Input
                          type="datetime-local"
                          name="departureTime"
                          min={dayjs().format("YYYY-MM-DDTHH:mm")}
                          required
                        />
                      </div>

                      <div className="space-y-2">
                        <Label htmlFor="flightCode">Mã chuyến bay</Label>
                        <Select name="flightCode" required>
                          <SelectTrigger>
                            <SelectValue placeholder="Chọn mã chuyến bay" />
                          </SelectTrigger>
                          <SelectContent>
                            {planes.map((plane) => (
                              <SelectItem key={plane.plane_code} value={plane.plane_code}>
                                {plane.plane_name}
                              </SelectItem>
                            ))}
                          </SelectContent>
                        </Select>
                      </div>

                      <div className="space-y-2">
                        <Label htmlFor="price">Giá chuyến bay (VND)</Label>
                        <Input
                          type="number"
                          name="price"
                          min={0}
                          required
                        />
                      </div>

                      <div className="space-y-4">
                        <h3 className="text-lg font-semibold">Sân bay trung gian</h3>
                        {stopAirports.map((stop, index) => (
                          <div key={stop.key} className="space-y-2 p-4 border rounded-lg">
                            <Select
                              value={stop.airport}
                              onValueChange={(value) => {
                                const updated = [...stopAirports];
                                updated[index].airport = value;
                                setStopAirports(updated);
                              }}
                            >
                              <SelectTrigger>
                                <SelectValue placeholder="Chọn sân bay trung gian" />
                              </SelectTrigger>
                              <SelectContent>
                                {airports.map((a) => (
                                  <SelectItem key={a.airport_code} value={a.airport_code}>
                                    {a.airport_name}
                                  </SelectItem>
                                ))}
                              </SelectContent>
                            </Select>

                            <Input
                              type="number"
                              placeholder="Thời gian dừng (phút)"
                              min={parameter.min_intermediate_stop_duration}
                              max={parameter.max_intermediate_stop_duration}
                              value={stop.time}
                              onChange={(e) => {
                                const updated = [...stopAirports];
                                updated[index].time = Number(e.target.value);
                                setStopAirports(updated);
                              }}
                            />

                            <Input
                              placeholder="Ghi chú"
                              value={stop.note}
                              onChange={(e) => {
                                const updated = [...stopAirports];
                                updated[index].note = e.target.value;
                                setStopAirports(updated);
                              }}
                            />

                            <Button
                              variant="destructive"
                              onClick={() => {
                                const updated = stopAirports.filter((_, i) => i !== index);
                                setStopAirports(updated);
                              }}
                            >
                              <IconTrash className="w-4 h-4 mr-2" />
                              Xóa
                            </Button>
                          </div>
                        ))}

                        {stopAirports.length < parameter.max_intermediate_stops && (
                          <Button
                            variant="outline"
                            onClick={addStopAirport}
                            className="w-full"
                          >
                            <IconPlus className="w-4 h-4 mr-2" />
                            Thêm sân bay trung gian
                          </Button>
                        )}
                      </div>

                      <Button type="submit" className="w-full">
                        Thêm chuyến bay
                      </Button>
                    </form>
                  </CardContent>
                </Card>
              </div>
            </div>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
} 