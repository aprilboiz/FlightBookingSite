import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api"

export const getAirports = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/airports`);
        return response.data;
    } catch (error) {
        console.error("Error fetching airports:", error);
        throw error;
    }
}

export const getAirportByCode = async (code) => {
    try {
        const response = await axios.get(`${API_BASE_URL}/airports/${code}`);
        return response.data;
    } catch (error) {
        console.error("Error fetching airport by code:", error);
        throw error;
    }
}
