import React, { useState } from "react";
import {
  Form,
  Input,
  Table,
  Row,
  Col,
  Typography,
  Select,
  InputNumber,
  Button,
  notification,
} from "antd";

const { Title } = Typography;
const { Option } = Select;

const airports = [
  { value: "HAN", label: "Hà Nội" },
  { value: "SGN", label: "Hồ Chí Minh" },
  { value: "DAD", label: "Đà Nẵng" },
  { value: "CXR", label: "Nha Trang" },
  { value: "HUI", label: "Huế" },
  { value: "VCA", label: "Cần Thơ" },
  { value: "PQC", label: "Phú Quốc" },
  { value: "VII", label: "Vinh" },
  { value: "HPH", label: "Hải Phòng" },
  { value: "THD", label: "Thanh Hóa" },
];

const columns = [
  { title: "STT", dataIndex: "stt", key: "stt", width: 50 },
  {
    title: "Sân Bay Trung Gian",
    dataIndex: "airport",
    key: "airport",
    render: () => (
      <Select defaultValue="" className="w-full">
        <Option value="">Chọn sân bay</Option>
        {airports.map((a) => (
          <Option key={a.value} value={a.value}>
            {a.label}
          </Option>
        ))}
      </Select>
    ),
  },
  {
    title: "Thời gian dừng",
    dataIndex: "time",
    key: "time",
    render: () => (
      <InputNumber min={10} max={20} defaultValue={10} className="w-full" />
    ),
  },
  {
    title: "Ghi chú",
    dataIndex: "note",
    key: "note",
    render: () => <Input className="w-full" />,
  },
];

const data = [
  { key: "1", stt: 1, airport: "", time: "", note: "" },
  { key: "2", stt: 2, airport: "", time: "", note: "" },
];

const CalenderPlane = () => {
  const [selectedDate, setSelectedDate] = React.useState(
    new Date().toISOString().split("T")[0]
  );

  const handleDateChange = (e) => {
    const chosenDate = e.target.value;
    const today = new Date().toISOString().split("T")[0];

    if (chosenDate < today) {
        notification.error({
            message: "Ngày không hợp lệ",
            description: "Ngày không thể nhỏ hơn ngày hiện tại",
        });
      setSelectedDate(today);
    } else {
      setSelectedDate(chosenDate);
    }
  };
  return (
    <div className="flex flex-col gap-5 w-1/2 mx-auto">
      <h2 className="text-center text-xl font-bold">Lịch chuyến bay</h2>
      <Form className="">
        <div>
          <Form.Item label="Mã chuyến bay">
            <Input value="RuaAirline-001" disabled />
          </Form.Item>
          <Form.Item label="Giá vé">
            <Input placeholder="Nhập giá vé" />
          </Form.Item>
        </div>

        <div>
          <Form.Item label="Sân bay đi">
            <Select placeholder="Chọn sân bay" className="w-full">
              {airports.map((airport) => (
                <Option key={airport.value} value={airport.value}>
                  {airport.label}
                </Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item label="Sân bay đến">
            <Select placeholder="Chọn sân bay" className="w-full">
              {airports.map((airport) => (
                <Option key={airport.value} value={airport.value}>
                  {airport.label}
                </Option>
              ))}
            </Select>
          </Form.Item>
        </div>

        <div>
          <Form.Item label="Ngày - Giờ">
            <Input
              type="date"
              value={selectedDate}
              onChange={handleDateChange}
              placeholder="Nhập ngày - giờ"
            />
          </Form.Item>
          <Form.Item label="Thời gian bay (phút)">
            <InputNumber min={30} defaultValue={30} className="w-full" />
          </Form.Item>
        </div>

        <div>
          <Form.Item label="Số lượng ghế hạng 1">
            <Input placeholder="Nhập số lượng" />
          </Form.Item>
          <Form.Item label="Số lượng ghế hạng 2">
            <Input placeholder="Nhập số lượng" />
          </Form.Item>
        </div>
        <Table
          columns={columns}
          dataSource={data}
          pagination={false}
          bordered
          className="mt-5"
        />
        <div className="flex justify-end mt-5">
          <Button type="primary" htmlType="submit">
            Thêm chuyến bay
          </Button>
        </div>
      </Form>
    </div>
  );
};

export default CalenderPlane;
