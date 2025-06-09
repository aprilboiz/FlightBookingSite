export interface Airport {
  airport_code: string;
  airport_name: string;
}

export interface Plane {
  plane_code: string;
  plane_name: string;
}

export interface Parameter {
  min_flight_duration: number;
  max_intermediate_stops: number;
  min_intermediate_stop_duration: number;
  max_intermediate_stop_duration: number;
}

export interface StopAirport {
  key: string;
  airport: string;
  time: number;
  note: string;
}

export interface FlightFormData {
  departureAirport: string;
  arrivalAirport: string;
  duration: number;
  departureTime: any; // dayjs object
  flightCode: string;
  price: number;
}

export interface FlightData {
  arrival_airport: string;
  base_price: number;
  departure_airport: string;
  departure_date: string;
  duration: number;
  intermediate_stops: {
    note: string;
    stop_airport: string;
    stop_duration: number;
    stop_order: number;
  }[];
  plane_code: string;
} 