export interface Seat {
  seat_number: string;
  class_name: string;
  price: number;
  is_booked: boolean;
}

export interface FlightDetails {
  flight_code: string;
  seat_class_info: {
    class_name: string;
  }[];
  seats: Seat[];
}

export interface TicketData {
  booking_type: string;
  email?: string;
  flight_code: string;
  full_name: string;
  id_card: string;
  phone_number: string;
  seat_number: string;
}

export interface BookingType {
  types: string[];
} 