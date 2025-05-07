import React, { useState, useEffect } from "react";
import { getFlights } from "../services/flightService";
// import { getAirports } from "../services/airportService.js";
import { useNavigate } from "react-router-dom";
import {
  Table,
  Button,
  // Modal,
  Form,
  // Input,
  // InputNumber,
  // DatePicker,
  // Space,
  notification,
  // Select,
} from "antd";
import {
  // EditOutlined,
  // DeleteOutlined,
  // PlusOutlined,
  // MinusCircleOutlined,
  InfoCircleOutlined,
} from "@ant-design/icons";
import dayjs from "dayjs";

const ListPlane = () => {
  const [flights, setFlights] = useState([]);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [form] = Form.useForm();
  const [editingFlight, setEditingFlight] = useState(null);

  const [airports, setAirports] = useState([]);

  const navigate = useNavigate();

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      const data = await getFlights();
      const formatted = data.map((item, index) => ({
        key: index,
        flightCode: item.flight_code,
        departureAirport: item.departure_airport,
        arrivalAirport: item.arrival_airport,
        departureTime: item.departure_date,
        price: item.base_price,
        empty_seat: item.empty_seats,
        booked_seat: item.booked_seats,
        total_seat: item.total_seats,
        intermediateStops: item.intermediate_stops || [],
      }));
      setFlights(formatted);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải danh sách chuyến bay",
      });
    }
  };


  const columns = [
    {
      title: "Mã chuyến bay",
      dataIndex: "flightCode",
      key: "flightCode",
    },
    {
      title: "Sân bay đi",
      dataIndex: "departureAirport",
      key: "departureAirport",
    },
    {
      title: "Sân bay đến",
      dataIndex: "arrivalAirport",
      key: "arrivalAirport",
    },
    {
      title: "Thời gian khởi hành",
      dataIndex: "departureTime",
      key: "departureTime",
      render: (text) => dayjs(text).format("YYYY-MM-DD HH:mm:ss"),
    },
    {
      title: "Giá vé",
      dataIndex: "price",
      key: "price",
    },
    {
      title: "Số ghế còn trống",
      dataIndex: "empty_seat",
      key: "empty_seat",
    },
    {
      title: "Số ghế đã đặt",
      dataIndex: "booked_seat",
      key: "booked_seat",
    },
    {
      title: "Tổng số ghế",
      dataIndex: "total_seat",
      key: "total_seat",
    },
    {
      title: "Hành động",
      key: "action",
      render: (_, record) => (
        <Button
          icon={<InfoCircleOutlined />}
          onClick={() => navigate(`/list-plane/${record.flightCode}`)}
        />
      ),
    },
  ];

  return (
    <>
      <div className="w-full p-5 flex flex-col gap-5 items-center">
        <h2>DANH SÁCH CHUYẾN BAY</h2>
        <div className="w-3/4">
          <Table dataSource={flights} columns={columns} />
        </div>
      </div>
    </>
  );
};

export default ListPlane;
