import axios from "axios";
import { MonthlyReportResponse, YearlyReportResponse } from "@/types/report";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";

const getAuthHeader = () => {
  const token = localStorage.getItem("token");
  return {
    Authorization: `Bearer ${token}`,
  };
};

export const getReportMonth = async (month: number, year: number): Promise<MonthlyReportResponse> => {
  try {
    const response = await axios.get(`${API_URL}/reports/revenue`, {
      params: {
        month,
        year,
      },
      headers: getAuthHeader(),
    });
    return response.data;
  } catch (error: any) {
    throw new Error(
      error.response?.data?.message ||
        "Có lỗi xảy ra khi tải báo cáo doanh thu theo tháng"
    );
  }
};

export const getReportYear = async (year: number): Promise<YearlyReportResponse> => {
  try {
    const response = await axios.get(`${API_URL}/reports/revenue/yearly`, {
      params: {
        year,
      },
      headers: getAuthHeader(),
    });
    return response.data;
  } catch (error: any) {
    throw new Error(
      error.response?.data?.message ||
        "Có lỗi xảy ra khi tải báo cáo doanh thu theo năm"
    );
  }
}; 