import axios from 'axios';
import { LoginRequest, LoginResponse, ApiError } from '@/types/auth';
import { Airport, Plane, Parameter, FlightData } from '@/types/flight';
import { FlightDetails, TicketData, BookingType } from '@/types/ticket';

const API_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor for adding auth token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Add response interceptor for handling errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const apiError: ApiError = {
      message: error.response?.data?.message || 'An error occurred',
      status: error.response?.status || 500,
    };
    return Promise.reject(apiError);
  }
);

export const authService = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    const response = await api.post<LoginResponse>('/auth/login', credentials);
    return response.data;
  },
};

export const airportService = {
  getAirports: async (): Promise<Airport[]> => {
    const response = await api.get<Airport[]>('/airports');
    return response.data;
  },
};

export const planeService = {
  getPlanes: async (): Promise<Plane[]> => {
    const response = await api.get<Plane[]>('/planes');
    return response.data;
  },
};

export const parameterService = {
  getParameters: async (): Promise<Parameter> => {
    const response = await api.get<Parameter>('/params');
    return response.data;
  },
};

export const flightService = {
  getFlights: async (): Promise<any[]> => {
    const response = await api.get('/flights');
    return response.data;
  },
  getFlightByCode: async (code: string): Promise<FlightDetails> => {
    const response = await api.get<FlightDetails>(`/flights/${code}`);
    return response.data;
  },
  addFlight: async (flightData: FlightData): Promise<void> => {
    await api.post('/flights', flightData);
  },
};

export const ticketService = {
  addTicket: async (ticketData: TicketData): Promise<void> => {
    await api.post('/tickets', ticketData);
  },
  getTickets: async (): Promise<any[]> => {
    const response = await api.get('/tickets');
    return response.data;
  },
  cancelTicket: async (ticketId: string): Promise<void> => {
    await api.delete(`/tickets/${ticketId}`);
  },
  getStatusTicket: async (): Promise<any> => {
    const response = await api.get('/tickets/statuses');
    return response.data;
  },
  updateStatusTicket: async (ticketId: string, status: string): Promise<void> => {
    await api.put(`/tickets/${ticketId}/status`, { status });
  },
  getBookingTypes: async (): Promise<BookingType> => {
    const response = await api.get<BookingType>('/tickets/booking-types');
    return response.data;
  },
};

export default api; 