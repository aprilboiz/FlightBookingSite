import React, { useState } from "react";
import {
  Table,
  Button,
  Modal,
  Form,
  Input,
  InputNumber,
  DatePicker,
  notification,
} from "antd";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";
import dayjs from "dayjs";

const ListPlane = () => {
  const [flights, setFlights] = useState([
    {
      key: "1",
      from: "Hà Nội",
      to: "TP. Hồ Chí Minh",
      via1: "Đà Nẵng",
      via2: "Huế",
      dateTime: "2025-04-10 08:00",
      duration: "150",
      price: 2000000,
      firstClassSeats: 10,
      secondClassSeats: 100,
    },
    {
      key: "2",
      from: "Đà Nẵng",
      to: "Phú Quốc",
      via1: "",
      via2: "",
      dateTime: "2025-04-12 14:00",
      duration: "180",
      price: 2500000,
      firstClassSeats: 8,
      secondClassSeats: 90,
    },
  ]);

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingFlight, setEditingFlight] = useState(null);
  const [form] = Form.useForm();

  const handleEdit = (record) => {
    setEditingFlight(record);
    form.setFieldsValue({
      ...record,
      dateTime: dayjs(record.dateTime, "YYYY-MM-DD HH:mm"),
    });
    setIsModalOpen(true);
  };

  const handleDelete = (key) => {
    // setFlights(flights.filter((flight) => flight.key !== key));
  };

  const handleOk = () => {
    form.validateFields().then((values) => {
      const selectedDate = values.dateTime;
      const now = dayjs();

      if (!selectedDate || selectedDate.isBefore(now)) {
        notification.error({
          message: "Ngày - Giờ phải lớn hơn thời gian hiện tại!",
        });
        return;
      }

      if (values.duration < 30) {
        notification.error({ message: "Thời gian bay phải ít nhất 30 phút!" });
        return;
      }

      setIsModalOpen(false);
    });
  };

  const columns = [
    {
      title: "Từ",
      dataIndex: "from",
      key: "from",
    },
    {
      title: "Đến",
      dataIndex: "to",
      key: "to",
    },
    {
      title: "Trung gian",
      key: "via",
      render: (_, record) =>
        [record.via1, record.via2].filter(Boolean).join(", ") || "-",
    },
    {
      title: "Ngày - Giờ",
      dataIndex: "dateTime",
      key: "dateTime",
    },
    {
      title: "Thời gian bay",
      dataIndex: "duration",
      key: "duration",
    },
    {
      title: "Giá vé",
      dataIndex: "price",
      key: "price",
      render: (price) => `${price.toLocaleString()} VND`,
    },
    {
      title: "Ghế hạng 1",
      dataIndex: "firstClassSeats",
      key: "firstClassSeats",
    },
    {
      title: "Ghế hạng 2",
      dataIndex: "secondClassSeats",
      key: "secondClassSeats",
    },
    {
      title: "Tổng số ghế",
      key: "totalSeats",
      render: (_, record) => record.firstClassSeats + record.secondClassSeats,
    },
    {
      title: "Hành động",
      key: "actions",
      render: (_, record) => (
        <>
          <Button
            icon={<EditOutlined />}
            onClick={() => handleEdit(record)}
            style={{ marginRight: 8 }}
          />
          <Button
            icon={<DeleteOutlined />}
            onClick={() => handleDelete(record.key)}
            danger
          />
        </>
      ),
    },
  ];

  return (
    <>
      <Table columns={columns} dataSource={flights} pagination={false} />
      <Modal
        title="Chỉnh sửa chuyến bay"
        open={isModalOpen}
        onOk={handleOk}
        onCancel={() => setIsModalOpen(false)}
      >
        <Form form={form} layout="vertical">
          <Form.Item
            name="from"
            label="Từ"
            rules={[{ required: true, message: "Vui lòng nhập nơi đi!" }]}
          >
            {" "}
            <Input />{" "}
          </Form.Item>
          <Form.Item
            name="to"
            label="Đến"
            rules={[{ required: true, message: "Vui lòng nhập nơi đến!" }]}
          >
            {" "}
            <Input />{" "}
          </Form.Item>
          <Form.Item name="via1" label="Trung gian 1">
            {" "}
            <Input />{" "}
          </Form.Item>
          <Form.Item name="via2" label="Trung gian 2">
            {" "}
            <Input />{" "}
          </Form.Item>
          <Form.Item
            name="dateTime"
            label="Ngày"
            rules={[{ required: true, message: "Vui lòng chọn ngày!" }]}
          >
            <DatePicker format="MM/DD/YYYY" style={{ width: "100%" }} />
          </Form.Item>
          <Form.Item
            name="duration"
            label="Thời gian bay (phút)"
            rules={[
              { required: true, message: "Vui lòng nhập thời gian bay!" },
            ]}
          >
            <InputNumber style={{ width: "100%" }} min={30} />
          </Form.Item>
          <Form.Item name="price" label="Giá vé" rules={[{ required: true }]}>
            {" "}
            <InputNumber style={{ width: "100%" }} />{" "}
          </Form.Item>
          <Form.Item
            name="firstClassSeats"
            label="Ghế hạng 1"
            rules={[{ required: true }]}
          >
            {" "}
            <InputNumber style={{ width: "100%" }} />{" "}
          </Form.Item>
          <Form.Item
            name="secondClassSeats"
            label="Ghế hạng 2"
            rules={[{ required: true }]}
          >
            {" "}
            <InputNumber style={{ width: "100%" }} />{" "}
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default ListPlane;
