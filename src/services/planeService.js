import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api"

export const getPlane = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/planes`);
        return response.data;
    } catch (error) {
        console.error("Error fetching plane:", error);
        throw error;
    }
}
