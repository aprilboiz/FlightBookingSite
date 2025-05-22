"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { toast } from "sonner";
import { IconInfoCircle } from "@tabler/icons-react";
import dayjs from "dayjs";
import { flightService } from "@/services/api";

interface Flight {
  key: number;
  flightCode: string;
  departureAirport: string;
  arrivalAirport: string;
  departureTime: string;
  price: number;
  empty_seat: number;
  booked_seat: number;
  total_seat: number;
  intermediateStops: any[];
}

export default function ListFlightPage() {
  const router = useRouter();
  const [flights, setFlights] = useState<Flight[]>([]);

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      const data = await flightService.getFlights();
      const formatted = data.map((item: any, index: number) => ({
        key: index,
        flightCode: item.flight_code,
        departureAirport: item.departure_airport,
        arrivalAirport: item.arrival_airport,
        departureTime: item.departure_date_time,
        price: item.base_price,
        empty_seat: item.empty_seats,
        booked_seat: item.booked_seats,
        total_seat: item.total_seats,
        intermediateStops: item.intermediate_stops || [],
      }));
      setFlights(formatted);
    } catch (error) {
      toast.error("Không thể tải danh sách chuyến bay");
    }
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
                    <CardTitle className="text-2xl text-center">DANH SÁCH CHUYẾN BAY</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="rounded-md border">
                      <Table>
                        <TableHeader>
                          <TableRow>
                            <TableHead>Mã chuyến bay</TableHead>
                            <TableHead>Sân bay đi</TableHead>
                            <TableHead>Sân bay đến</TableHead>
                            <TableHead>Thời gian khởi hành</TableHead>
                            <TableHead>Giá vé</TableHead>
                            <TableHead>Số ghế còn trống</TableHead>
                            <TableHead>Số ghế đã đặt</TableHead>
                            <TableHead>Tổng số ghế</TableHead>
                          </TableRow>
                        </TableHeader>
                        <TableBody>
                          {flights.map((flight) => (
                            <TableRow key={flight.key}>
                              <TableCell>{flight.flightCode}</TableCell>
                              <TableCell>{flight.departureAirport}</TableCell>
                              <TableCell>{flight.arrivalAirport}</TableCell>
                              <TableCell>
                                {dayjs(flight.departureTime).format("YYYY-MM-DD HH:mm:ss")}
                              </TableCell>
                              <TableCell>{flight.price.toLocaleString()} VND</TableCell>
                              <TableCell>{flight.empty_seat}</TableCell>
                              <TableCell>{flight.booked_seat}</TableCell>
                              <TableCell>{flight.total_seat}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </div>
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