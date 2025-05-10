import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api"

export const getReportMonth = async (month, year) => {
    try {
        const response = await axios.get(`${API_BASE_URL}/reports/revenue`, {
            params: {
                month,
                year
            }
        });
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi tải báo cáo doanh thu theo tháng');
    }
}

export const getReportYear = async (year) => {
    try {
        const response = await axios.get(`${API_BASE_URL}/reports/revenue/yearly`, {
            params: {
                year
            }
        });
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi tải báo cáo doanh thu theo năm');
    }
}