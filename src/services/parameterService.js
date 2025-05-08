import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api"

export const getParameter = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/params`);
        return response.data;
    } catch (error) {
        console.error("Error fetching parameters:", error);
    }
}

export const updateParameter = async (params) => {
    try {
        const response = await axios.put(`${API_BASE_URL}/params`, params, {
            headers: {
                "Content-Type": "application/json",
            },
        });
        return response.data;
    } catch (error) {
        console.error("Error updating parameters:", error);
    }
}