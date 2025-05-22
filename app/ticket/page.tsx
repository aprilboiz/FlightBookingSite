"use client";

import { useState, useEffect } from "react";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Input } from "@/components/ui/input";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { toast } from "sonner";
import { flightService } from "@/services/api";
import { TicketBooking } from "@/components/ticket-booking";
import dayjs from "dayjs";

interface Flight {
  key: number;
  flight_code: string;
  departure_airport: string;
  arrival_airport: string;
  departure_date_time: string;
}

export default function TicketPage() {
  const [flights, setFlights] = useState<Flight[]>([]);
  const [filteredFlights, setFilteredFlights] = useState<Flight[]>([]);
  const [selectedFlight, setSelectedFlight] = useState<Flight | null>(null);

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      const data = await flightService.getFlights();
      const formatted = data.map((item: any, index: number) => ({
        key: index,
        flight_code: item.flight_code,
        departure_airport: item.departure_airport,
        arrival_airport: item.arrival_airport,
        departure_date_time: item.departure_date_time,
      }));
      setFlights(formatted);
      setFilteredFlights(formatted);
    } catch (error) {
      toast.error("Không thể tải danh sách chuyến bay");
    }
  };

  const handleSearch = (value: string) => {
    const result = flights.filter((flight) =>
      flight.flight_code.toLowerCase().includes(value.toLowerCase())
    );
    setFilteredFlights(result);
  };

  const handleBookingSuccess = () => {
    setSelectedFlight(null);
    fetchFlights(); // Refresh flight list to update seat availability
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
                <div className="flex justify-between items-start gap-5">
                  <div className="w-3/4">
                    <Input
                      placeholder="Tìm kiếm mã chuyến bay"
                      onChange={(e) => handleSearch(e.target.value)}
                      className="mb-4"
                    />

                    <Table>
                      <TableHeader>
                        <TableRow>
                          <TableHead>Mã chuyến bay</TableHead>
                          <TableHead>Điểm đi</TableHead>
                          <TableHead>Điểm đến</TableHead>
                          <TableHead>Thời gian</TableHead>
                        </TableRow>
                      </TableHeader>
                      <TableBody>
                        {filteredFlights.map((flight) => (
                          <TableRow
                            key={flight.flight_code}
                            className="cursor-pointer hover:bg-muted/50"
                            onClick={() => setSelectedFlight(flight)}
                          >
                            <TableCell>{flight.flight_code}</TableCell>
                            <TableCell>{flight.departure_airport}</TableCell>
                            <TableCell>{flight.arrival_airport}</TableCell>
                            <TableCell>
                              {dayjs(flight.departure_date_time).format("YYYY-MM-DD HH:mm:ss")}
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </div>
                  <div>
                    <TicketBooking 
                      selectedFlight={selectedFlight} 
                      onSuccess={handleBookingSuccess}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
} 