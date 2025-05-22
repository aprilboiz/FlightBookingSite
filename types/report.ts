export interface FlightReport {
  flightCode: string;
  tickets: number;
  revenue: number;
}

export interface MonthReport {
  month: number;
  flightCount: number;
  revenue: number;
}

export interface MonthlyReportResponse {
  month: number;
  year: number;
  flights: FlightReport[];
  totalRevenue: number;
  totalTickets: number;
}

export interface YearlyReportResponse {
  year: number;
  months: MonthReport[];
  totalRevenue: number;
  totalFlights: number;
} 