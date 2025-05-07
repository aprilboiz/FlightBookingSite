import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getAirports } from "../services/airportService.js";
import {Form, Input, Typography, Select, InputNumber, Button, notification, DatePicker, Space, message,} from "antd";
import dayjs from "dayjs";
import { PlusOutlined } from "@ant-design/icons";
import { getPlane } from "../services/planeService.js";
import { addFlight } from "../services/flightService.js";
import { getParameter } from "../services/parameterService.js";

const { Title } = Typography;
const { Option } = Select;

const CalenderPlane = () => {
  const [airports, setAirports] = useState([]);
  const [planes, setPlanes] = useState([]);
  const [stopAirports, setStopAirports] = useState([]);
  const [parameter, setParameter] = useState([]);

  const navigate = useNavigate();

  const fetchAirports = async () => {
    try {
      const data = await getAirports();
      setAirports(data);
    } catch (err) {
      notification.error({
        message: "Lỗi",
        description: "Không thể lấy danh sách sân bay",
      });
    }
  };

  const fetchPlane = async () => {
    try {
      const data = await getPlane();
      setPlanes(data);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể lấy danh sách máy bay",
      });
    }
  };

  const fetchParameter = async () => {
    try {
      const data = await getParameter();
      setParameter(data);
      console.log("Parameter data:", parameter);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể lấy danh sách tham số",
      });
    }
  };

  useEffect(() => {
    fetchAirports();
    fetchPlane();
    fetchParameter();
  }, []);

  const onFinish = async (values) => {
    const { departureAirport, arrivalAirport, duration, departureTime, flightCode, price } = values;

    if (departureAirport === arrivalAirport) {
      notification.error({
        message: "Lỗi",
        description: "Sân bay đi và đến không được trùng nhau",
      });
      return;
    }

    const middleAirports = stopAirports.map((a) => a.airport).filter(Boolean);

    const hasInvalidMiddle = middleAirports.some(
      (a) => a === departureAirport || a === arrivalAirport
    );

    if (hasInvalidMiddle) {
      notification.error({
        message: "Lỗi",
        description: "Sân bay trung gian không được trùng sân bay đi hoặc đến",
      });
      return;
    }

    const hasDuplicateMiddle =
      new Set(middleAirports).size !== middleAirports.length;
    if (hasDuplicateMiddle) {
      notification.error({
        message: "Lỗi",
        description: "Không được chọn 2 sân bay trung gian giống nhau",
      });
      return;
    }

    const hasInvalidStopTime = stopAirports.some(
      (stop) => stop.time < parameter.min_intermediate_stop_duration || stop.time > parameter.max_intermediate_stop_duration
    );
    if (hasInvalidStopTime) {
      notification.error({
        message: "Lỗi",
        description: "Thời gian dừng phải từ 10 đến 20 phút",
      });
      return;
    }

    const formattedDepartureTime = departureTime.format("YYYY-MM-DD HH:mm:ss");

    const flightData = {
      arrival_airport: arrivalAirport,
      base_price: price,
      departure_airport: departureAirport,
      departure_date: formattedDepartureTime,
      duration: duration,
      intermediate_stops: stopAirports.map((stop, index) => ({
        note: stop.note,
        stop_airport: stop.airport,
        stop_duration: stop.time,
        stop_order: index + 1,
      })),
      plane_code: flightCode,
    };

    try {
      await addFlight(flightData);
      notification.success({
        message: "Thành công",
        description: "Chuyến bay đã được thêm thành công",
      });
      navigate("/list-plane")
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: error.message || "Có lỗi xảy ra khi thêm chuyến bay",
      });
    }

    console.log("Dữ liệu chuyến bay:", values, stopAirports);
    notification.success({
      message: "Thành công",
      description: "Chuyến bay đã được thêm",
    });
  };

  const addStopAirport = () => {
    if (stopAirports.length >= parameter.max_intermediate_stops) return;
    setStopAirports([
      ...stopAirports,
      { key: Date.now().toString(), airport: "", time: 0, note: "" },
    ]);
  };

  return (
    <div className="flex flex-col gap-5 w-2/3 mx-auto my-10">
      <div className="flex items-center gap-5">
        <p>Danh sách chuyến bay</p>
        <p>Thêm chuyến bay</p>
      </div>

      <Form
        layout="vertical"
        className="flex flex-col gap-5"
        onFinish={onFinish}
      >
        <Title level={3} className="text-center p-5">
          THÊM LỊCH CHUYẾN BAY
        </Title>

        <Form.Item
          name="departureAirport"
          label="Sân bay đi"
          rules={[{ required: true, message: "Vui lòng chọn sân bay đi" }]}
        >
          <Select placeholder="Chọn sân bay">
            {airports.map((airport) => (
              <Option key={airport.airport_code} value={airport.airport_code}>
                {airport.airport_name}
              </Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          name="arrivalAirport"
          label="Sân bay đến"
          rules={[{ required: true, message: "Vui lòng chọn sân bay đến" }]}
        >
          <Select placeholder="Chọn sân bay">
            {airports.map((airport) => (
              <Option key={airport.airport_code} value={airport.airport_code}>
                {airport.airport_name}
              </Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          name="duration"
          label="Thời gian bay (phút)"
          rules={[
            { required: true, message: "Vui lòng nhập thời gian bay" },
            {
              validator: (_, value) =>
                value >= parameter.min_flight_duration
                  ? Promise.resolve()
                  : Promise.reject(`Thời gian bay phải lớn hơn ${parameter.min_flight_duration} phút`),
            },
          ]}
        >
          <InputNumber className="w-full" min={parameter.min_flight_duration} />
        </Form.Item>

        <Form.Item
          name="departureTime"
          label="Ngày khởi hành"
          rules={[
            { required: true, message: "Vui lòng chọn ngày khởi hành" },
            {
              validator: (_, value) =>
                !value || value.isAfter(dayjs())
                  ? Promise.resolve()
                  : Promise.reject("Ngày giờ khởi hành phải lớn hơn hiện tại"),
            },
          ]}
        >
          <DatePicker showTime className="w-full" />
        </Form.Item>

        <Form.Item
          name="flightCode"
          label="Mã chuyến bay"
          rules={[{ required: true, message: "Vui lòng chọn mã chuyến bay" }]}
        >
          <Select placeholder="Chọn mã chuyến bay">
            {planes.map((plane) => (
              <Option key={plane.plane_code} value={plane.plane_code}>
                {plane.plane_name}
              </Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          name="price"
          label="Giá chuyến bay (VND)"
          rules={[{ required: true, message: "Vui lòng nhập giá" }]}
        >
          <InputNumber className="w-full" min={0} />
        </Form.Item>

        <Title level={4}>Sân bay trung gian</Title>
        {stopAirports.map((stop, index) => (
          <Space key={stop.key} direction="vertical" className="w-full">
            <Select
              placeholder="Chọn sân bay trung gian"
              className="w-full"
              value={stop.airport}
              onChange={(value) => {
                const updated = [...stopAirports];
                updated[index].airport = value;
                setStopAirports(updated);
              }}
            >
              {airports.map((a) => (
                <Option key={a.airport_code} value={a.airport_code}>
                  {a.airport_name}
                </Option>
              ))}
            </Select>
            <InputNumber
              placeholder="Thời gian dừng (phút)"
              className="w-full"
              min={10}
              max={20}
              value={stop.time}
              onChange={(value) => {
                const updated = [...stopAirports];
                updated[index].time = value;
                setStopAirports(updated);
              }}
            />
            <Input
              placeholder="Ghi chú"
              value={stop.note}
              onChange={(e) => {
                const updated = [...stopAirports];
                updated[index].note = e.target.value;
                setStopAirports(updated);
              }}
            />
          </Space>
        ))}

        {stopAirports.length < 2 && (
          <Button
            icon={<PlusOutlined />}
            onClick={addStopAirport}
            type="dashed"
            className="w-full"
          >
            Thêm sân bay trung gian
          </Button>
        )}

        <Form.Item>
          <Button type="primary" htmlType="submit" className="w-full">
            Thêm chuyến bay
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default CalenderPlane;
