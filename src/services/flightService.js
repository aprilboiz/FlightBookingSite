import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api"

export const addFlight = async (flightData) => {
    console.log("Flight data:", flightData); 
    try {
        const response = await axios.post(`${API_BASE_URL}/flights`, flightData, {
            headers: {
                "Content-Type": "application/json",
            },
        });
        return response.data;
    } catch (error) {
        throw new Error(error.response?.data?.message || 'Có lỗi xảy ra khi thêm chuyến bay');
    }
}