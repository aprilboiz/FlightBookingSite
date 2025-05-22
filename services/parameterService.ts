import axios from "axios";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";

const getAuthHeader = () => {
  const token = localStorage.getItem("token");
  return {
    Authorization: `Bearer ${token}`,
  };
};

export interface FlightParameter {
  number_of_airports: number;
  min_flight_duration: number;
  max_intermediate_stops: number;
  min_intermediate_stop_duration: number;
  max_intermediate_stop_duration: number;
  max_ticket_classes: number;
  latest_ticket_purchase_time: number;
  ticket_cancellation_time: number;
}

export const getParameter = async (): Promise<FlightParameter> => {
  try {
    const response = await axios.get(`${API_URL}/params`, {
      headers: getAuthHeader(),
    });
    return response.data;
  } catch (error: any) {
    throw new Error(
      error.response?.data?.message || "Có lỗi xảy ra khi tải quy định"
    );
  }
};

export const updateParameter = async (data: FlightParameter): Promise<void> => {
  try {
    await axios.put(`${API_URL}/params`, data, {
      headers: getAuthHeader(),
    });
  } catch (error: any) {
    throw new Error(
      error.response?.data?.message || "Có lỗi xảy ra khi cập nhật quy định"
    );
  }
}; 