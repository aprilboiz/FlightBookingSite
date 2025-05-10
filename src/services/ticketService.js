import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api"

export const addTicket = async (ticketData) => {
    try {
        const response = await axios.post(`${API_BASE_URL}/tickets`, ticketData, {
            headers: {
                "Content-Type": "application/json",
            },
        });
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.details || 'Có lỗi xảy ra khi thêm vé máy bay');
    }
}

export const getTickets = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/tickets`);
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi tải danh sách vé máy bay');
    }
}

export const cancelTicket = async (ticketId) => {
    try {
        const response = await axios.delete(`${API_BASE_URL}/tickets/${ticketId}`);
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi hủy vé máy bay');
    }
}

export const getStatusTicket = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/tickets/statuses`);
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi tải trạng thái vé máy bay');   
    }
}

export const updateStatusTicket = async (ticketId, status) => {
    try {
        const response = await axios.put(`${API_BASE_URL}/tickets/${ticketId}/status`, { status });
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi cập nhật trạng thái vé máy bay');
    }
}

export const getBookingTypes = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/tickets/booking-types`);
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi tải loại đặt vé máy bay');
    }
}


