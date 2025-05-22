"use client";

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { toast } from "sonner";
import { flightService, ticketService } from "@/services/api";
import { FlightDetails, Seat } from "@/types/ticket";

interface TicketBookingProps {
  selectedFlight: {
    flight_code: string;
    departure_airport: string;
    arrival_airport: string;
    departure_date_time: string;
  } | null;
  onSuccess?: () => void;
}

export function TicketBooking({ selectedFlight, onSuccess }: TicketBookingProps) {
  const [ticketPrice, setTicketPrice] = useState("0 VND");
  const [ticketClasses, setTicketClasses] = useState<string[]>([]);
  const [ticketTypes, setTicketTypes] = useState<string[]>([]);
  const [seats, setSeats] = useState<Seat[]>([]);
  const [formData, setFormData] = useState({
    flightCode: "",
    ticketType: "",
    passenger: "",
    idCard: "",
    phone: "",
    email: "",
    ticketClass: "",
    seat_number: "",
  });

  useEffect(() => {
    if (selectedFlight) {
      setFormData(prev => ({ ...prev, flightCode: selectedFlight.flight_code }));
      getFlightDetails(selectedFlight.flight_code);
      getTicketTypes();
    } else {
      setFormData({
        flightCode: "",
        ticketType: "",
        passenger: "",
        idCard: "",
        phone: "",
        email: "",
        ticketClass: "",
        seat_number: "",
      });
      setTicketPrice("0 VND");
      setTicketClasses([]);
      setTicketTypes([]);
      setSeats([]);
    }
  }, [selectedFlight]);

  const getFlightDetails = async (flightCode: string) => {
    try {
      const data = await flightService.getFlightByCode(flightCode);
      const classes = data.seat_class_info.map((item: any) => item.class_name);
      setTicketClasses(classes);
      setSeats(data.seats);
      if (classes.length > 0) {
        setFormData(prev => ({ ...prev, ticketClass: classes[0] }));
        const defaultSeat = data.seats.find((seat: any) => seat.class_name === classes[0]);
        setTicketPrice(defaultSeat ? `${defaultSeat.price} VND` : "0 VND");
      }
    } catch (error) {
      toast.error("Không thể tải thông tin máy bay");
    }
  };

  const getTicketTypes = async () => {
    try {
      const data = await ticketService.getBookingTypes();
      setTicketTypes(data.types);
    } catch (error) {
      toast.error("Không thể tải loại vé");
    }
  };

  const calculatePrice = (ticketClass: string) => {
    const seat = seats.find((seat) => seat.class_name === ticketClass);
    setTicketPrice(seat ? `${seat.price} VND` : "0 VND");
  };

  const handleClassChange = (value: string) => {
    setFormData(prev => ({ ...prev, ticketClass: value, seat_number: "" }));
    calculatePrice(value);
  };

  const handleSeatChange = (seatNumber: string) => {
    const seat = seats.find((seat) => seat.seat_number === seatNumber);
    setTicketPrice(seat ? `${seat.price} VND` : "0 VND");
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedFlight || !formData.seat_number) {
      toast.error("Vui lòng chọn chuyến bay và ghế");
      return;
    }

    const ticketData = {
      booking_type: formData.ticketType,
      email: formData.email || "",
      flight_code: selectedFlight.flight_code,
      full_name: formData.passenger,
      id_card: formData.idCard,
      phone_number: formData.phone,
      seat_number: formData.seat_number,
    };

    try {
      await ticketService.addTicket(ticketData);
      toast.success("Đặt vé thành công");
      setFormData({
        flightCode: "",
        ticketType: "",
        passenger: "",
        idCard: "",
        phone: "",
        email: "",
        ticketClass: "",
        seat_number: "",
      });
      setTicketPrice("0 VND");
      setTicketClasses([]);
      setSeats([]);
      onSuccess?.();
    } catch (error: any) {
      toast.error(error.message || "Có lỗi xảy ra khi đặt vé");
    }
  };

  return (
    <Card className="w-[500px]">
      <CardHeader>
        <CardTitle>Vé Chuyến Bay</CardTitle>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="flightCode">Chuyến bay</Label>
            <Input
              id="flightCode"
              value={formData.flightCode}
              disabled
              placeholder="Chọn chuyến bay"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="ticketType">Loại vé</Label>
            <Select
              value={formData.ticketType}
              onValueChange={(value) => setFormData(prev => ({ ...prev, ticketType: value }))}
              disabled={!selectedFlight}
            >
              <SelectTrigger>
                <SelectValue placeholder="Chọn loại vé" />
              </SelectTrigger>
              <SelectContent>
                {ticketTypes.map((type) => (
                  <SelectItem key={type} value={type}>
                    {type}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label htmlFor="passenger">Hành khách</Label>
            <Input
              id="passenger"
              value={formData.passenger}
              onChange={(e) => setFormData(prev => ({ ...prev, passenger: e.target.value }))}
              placeholder="Nhập tên hành khách"
              required
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="idCard">CMND</Label>
            <Input
              id="idCard"
              value={formData.idCard}
              onChange={(e) => setFormData(prev => ({ ...prev, idCard: e.target.value }))}
              placeholder="Nhập CMND"
              required
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="phone">Điện thoại</Label>
            <Input
              id="phone"
              value={formData.phone}
              onChange={(e) => setFormData(prev => ({ ...prev, phone: e.target.value }))}
              placeholder="Nhập số điện thoại"
              required
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              type="email"
              value={formData.email}
              onChange={(e) => setFormData(prev => ({ ...prev, email: e.target.value }))}
              placeholder="Nhập email (tùy chọn)"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="ticketClass">Hạng vé</Label>
            <Select
              value={formData.ticketClass}
              onValueChange={handleClassChange}
              disabled={!selectedFlight}
            >
              <SelectTrigger>
                <SelectValue placeholder="Chọn hạng vé" />
              </SelectTrigger>
              <SelectContent>
                {ticketClasses.map((cls) => (
                  <SelectItem key={cls} value={cls}>
                    {cls}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label htmlFor="seat_number">Số ghế</Label>
            <Select
              value={formData.seat_number}
              onValueChange={(value) => {
                setFormData(prev => ({ ...prev, seat_number: value }));
                handleSeatChange(value);
              }}
              disabled={!selectedFlight || !formData.ticketClass}
            >
              <SelectTrigger>
                <SelectValue placeholder="Chọn số ghế" />
              </SelectTrigger>
              <SelectContent>
                {seats
                  .filter((seat) => seat.class_name === formData.ticketClass)
                  .map((seat) => (
                    <SelectItem
                      key={seat.seat_number}
                      value={seat.seat_number}
                      disabled={seat.is_booked}
                    >
                      {seat.seat_number} {seat.is_booked ? "(Đã đặt)" : ""}
                    </SelectItem>
                  ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-2">
            <Label>Giá tiền</Label>
            <Input value={ticketPrice} disabled />
          </div>

          <Button
            type="submit"
            className="w-full"
            disabled={!selectedFlight}
          >
            Bán vé
          </Button>
        </form>
      </CardContent>
    </Card>
  );
} 