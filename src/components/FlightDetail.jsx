import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { notification, Descriptions, Spin } from "antd";
import dayjs from "dayjs";

import { getFlightByCode } from "../services/flightService";

const FlightDetail = () => {
  const { code } = useParams();
  const [flight, setFlight] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchFlightDetail();
  }, []);

  const fetchFlightDetail = async () => {
    try {
      const data = await getFlightByCode(code);
      setFlight(data);
      setLoading(false);
    } catch (error) {
      setLoading(false);
      notification.error({
        message: "Lỗi",
        description: "Không thể lấy thông tin chuyến bay",
      });
    }
  };

  if (loading) return <Spin size="large" className="m-10 block" />;

  return (
    <div className="p-5">
      <h2 className="text-xl font-bold mb-4">Chi tiết chuyến bay: {code}</h2>
      <Descriptions bordered column={1}>
        <Descriptions.Item label="Sân bay đi">
          {flight.departure_airport}
        </Descriptions.Item>
        <Descriptions.Item label="Sân bay đến">
          {flight.arrival_airport}
        </Descriptions.Item>
        <Descriptions.Item label="Thời gian khởi hành">
          {dayjs(flight.departure_date).format("YYYY-MM-DD HH:mm")}
        </Descriptions.Item>
        <Descriptions.Item label="Giá vé">
          {flight.base_price}
        </Descriptions.Item>
        <Descriptions.Item label="Mã máy bay">
          {flight.plane_code}
        </Descriptions.Item>
        <Descriptions.Item label="Thời lượng">
          {flight.duration} phút
        </Descriptions.Item>
        <Descriptions.Item label="Sân bay trung gian">
          {flight.intermediate_stops.length === 0 ? (
            "Không có"
          ) : (
            <ul>
              {flight.intermediate_stops.map((stop, index) => (
                <li key={index}>
                  <strong>{stop.stop_airport}</strong> - {stop.stop_duration}{" "}
                  phút
                  {stop.note && ` (${stop.note})`}
                </li>
              ))}
            </ul>
          )}
        </Descriptions.Item>
      </Descriptions>
    </div>
  );
};

export default FlightDetail;
