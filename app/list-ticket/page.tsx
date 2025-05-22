"use client";

import { useState, useEffect } from "react";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { toast } from "sonner";
import { flightService, ticketService } from "@/services/api";
import dayjs from "dayjs";

interface Ticket {
  key: number;
  ticketCode: string;
  flightCode: string;
  passengerName: string;
  phoneNumber: string;
  email: string;
  idCard: string;
  ticketPrice: string;
  paymentStatus: string;
  ticketType: string;
  departureDate: string;
}

export default function ListTicketPage() {
  const [tickets, setTickets] = useState<Ticket[]>([]);

  useEffect(() => {
    fetchTickets();
  }, []);

  const fetchTickets = async () => {
    try {
      const data = await ticketService.getTickets();
      const formatted = await Promise.all(
        data.map(async (item: any, index: number) => {
          const flight = await flightService.getFlightByCode(item.flight_code);
          return {
            key: index,
            ticketCode: item.id,
            flightCode: item.flight_code,
            passengerName: item.full_name,
            phoneNumber: item.phone_number,
            email: item.email,
            idCard: item.id_card,
            ticketPrice: item.price + " VNĐ",
            paymentStatus: item.ticket_status,
            ticketType: item.booking_type,
            departureDate: flight.departure_date_time
              ? dayjs(flight.departure_date_time).format("YYYY-MM-DD HH:mm:ss")
              : "Không có dữ liệu",
          };
        })
      );
      setTickets(formatted);
    } catch (error) {
      toast.error("Không thể tải danh sách vé máy bay");
    }
  };

  const handleUpdateStatus = async (ticketCode: string, departureDate: string) => {
    const currentDate = new Date();
    const flightDate = new Date(departureDate);
    if (flightDate <= currentDate) {
      try {
        await ticketService.updateStatusTicket(ticketCode, "EXPIRED");
        toast.success("Cập nhật trạng thái thành công");
        fetchTickets();
      } catch (error) {
        toast.error("Không thể cập nhật trạng thái vé");
      }
    }
  };

  const handleCancelTicket = async (ticketCode: string) => {
    try {
      await ticketService.updateStatusTicket(ticketCode, "CANCELLED");
      toast.success("Hủy vé thành công");
      fetchTickets();
    } catch (error) {
      toast.error("Không thể hủy vé");
    }
  };

  const getTicketTypeText = (type: string) => {
    if (type === "TICKET") return "Vé mua trực tiếp";
    if (type === "PLACE_ORDER") return "Vé đặt trước";
    return "Không xác định";
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
                <div className="w-full flex flex-col gap-5 items-center">
                  <h2 className="text-2xl font-bold">DANH SÁCH VÉ MÁY BAY</h2>
                  <div className="w-[80%]">
                    <Table>
                      <TableHeader>
                        <TableRow>
                          <TableHead>Mã vé</TableHead>
                          <TableHead>Mã chuyến bay</TableHead>
                          <TableHead>Tên hành khách</TableHead>
                          <TableHead>Số điện thoại</TableHead>
                          <TableHead>Email</TableHead>
                          <TableHead>CMND/CCCD</TableHead>
                          <TableHead>Giá vé</TableHead>
                          <TableHead>Trạng thái vé</TableHead>
                          <TableHead>Ngày khởi hành</TableHead>
                          <TableHead>Loại vé</TableHead>
                          <TableHead>Thao tác</TableHead>
                        </TableRow>
                      </TableHeader>
                      <TableBody>
                        {tickets.map((ticket) => (
                          <TableRow key={ticket.ticketCode}>
                            <TableCell>{ticket.ticketCode}</TableCell>
                            <TableCell>{ticket.flightCode}</TableCell>
                            <TableCell>{ticket.passengerName}</TableCell>
                            <TableCell>{ticket.phoneNumber}</TableCell>
                            <TableCell>{ticket.email}</TableCell>
                            <TableCell>{ticket.idCard}</TableCell>
                            <TableCell>{ticket.ticketPrice}</TableCell>
                            <TableCell>{ticket.paymentStatus}</TableCell>
                            <TableCell>{ticket.departureDate}</TableCell>
                            <TableCell>{getTicketTypeText(ticket.ticketType)}</TableCell>
                            <TableCell>
                              <div className="flex gap-2">
                                <Button
                                  onClick={() => handleCancelTicket(ticket.ticketCode)}
                                  variant="destructive"
                                  disabled={ticket.paymentStatus === "CANCELLED"}
                                >
                                  Hủy vé
                                </Button>
                              </div>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
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