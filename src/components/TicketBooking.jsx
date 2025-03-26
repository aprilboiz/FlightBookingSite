import React, { useState } from "react";
import { Form, Input, Select, Card, Button } from "antd";

const { Option } = Select;

const TicketBooking = () => {
  const [form] = Form.useForm();
  const [ticketPrice, setTicketPrice] = useState("1000 VND"); // Giá vé mặc định

  const handleClassChange = (value) => {
    const basePrice = 1000; // Giá vé hạng 2
    const price = value === "1" ? basePrice * 1.05 : basePrice;
    setTicketPrice(`${price} VND`);
  };
  return (
    <Card title="Vé Chuyến Bay" style={{ width: 500, margin: "auto" }}>
      <Form form={form} layout="vertical">
        <Form.Item label="Chuyến bay">
          <Input placeholder="Chuyến bay" disabled />
        </Form.Item>

        <Form.Item label="Hành khách" name="passenger">
          <Input placeholder="Nhập tên hành khách" />
        </Form.Item>

        <Form.Item label="CMND" name="idCard">
          <Input placeholder="Nhập CMND" />
        </Form.Item>

        <Form.Item label="Điện thoại" name="phone">
          <Input placeholder="Nhập số điện thoại" />
        </Form.Item>

        <Form.Item label="Hạng vé" name="ticketClass">
          <Select placeholder="Chọn hạng vé" onChange={handleClassChange}>
            <Option value="1">Hạng 1</Option>
            <Option value="2">Hạng 2</Option>
          </Select>
        </Form.Item>

        <Form.Item label="Giá tiền">
          <Input value={ticketPrice} disabled />
        </Form.Item>

        <Button type="primary" className="w-full" htmlType="submit">
          Bán vé
        </Button>
      </Form>
    </Card>
  );
};

export default TicketBooking;
