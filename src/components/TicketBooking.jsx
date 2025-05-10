import React, { useState, useEffect } from "react";
import { Form, Input, Select, Card, Button } from "antd";
import { getFlightByCode } from "../services/flightService.js";
import { addTicket, getBookingTypes } from "../services/ticketService.js";
import { notification } from "antd";


const { Option } = Select;

const TicketBooking = ({ selectedFlight }) => {
  const [form] = Form.useForm();
  const [ticketPrice, setTicketPrice] = useState("0 VND");
  const [ticketClasses, setTicketClasses] = useState([]);
  const [ticketTypes, setTicketTypes] = useState([]);
  const [seats, setSeats] = useState([]);
  const [selectedClass, setSelectedClass] = useState(null);

  useEffect(() => {
    if (selectedFlight){
      console.log("Selected flight:", selectedFlight);
      form.setFieldsValue(
        {
          "flightCode": selectedFlight.flight_code,
        }
      )
      getFlightDetails(selectedFlight.flight_code);
      getTicketTypes();
    } else {
      form.resetFields();
      setTicketPrice("0 VND");
      setTicketClasses([]);
      setTicketTypes([]);
      setSeats([]);
      setSelectedClass(null);
    }
  }, [selectedFlight, form])

  const getFlightDetails = async (flightCode) => {
    try {
      const data = await getFlightByCode(flightCode);
      const classes = data.seat_class_info.map((item) => item.class_name);
      setTicketClasses(classes);
      setSeats(data.seats);
      if (classes.length > 0) {
        form.setFieldsValue({ ticketClass: classes[0] });
        setSelectedClass(classes[0]);
        const defaultSeat = data.seats.find((seat) => seat.class_name === classes[0]);
        setTicketPrice(defaultSeat ? `${defaultSeat.price} VND` : "0 VND");
      }
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải thông tin máy bay",
      });
    }
  }

  const getTicketTypes = async () => {
    try {
      const data = await getBookingTypes();
      setTicketTypes(data.types);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải loại vé",
      });
    }
  };

  const calculatePrice = (ticketClass) => {
    const seat = seats.find((seat) => seat.class_name === ticketClass);
    setTicketPrice(seat ? `${seat.price} VND` : "0 VND");
  };

  const handleClassChange = (value) => {
    setSelectedClass(value);
    calculatePrice(value);
    form.setFieldsValue({ seat_number: undefined });
  };

  const handleSeatChange = (seatNumber) => {
    const seat = seats.find((seat) => seat.seat_number === seatNumber);
    setTicketPrice(seat ? `${seat.price} VND` : "0 VND");
  };

  const handleSubmit = async (values) => {
    if (!selectedFlight || !values.seat_number) {
      notification.error({
        message: "Lỗi",
        description: "Vui lòng chọn chuyến bay và ghế",
      });
      return;
    }

    const ticketData = {
      booking_type: values.ticketType,
      email: values.email || "",
      flight_code: selectedFlight.flight_code,
      full_name: values.passenger,
      id_card: values.idCard,
      phone_number: values.phone,
      seat_number: values.seat_number,
    };
    try {
      await addTicket(ticketData);
      notification.success({
        message: "Thành công",
        description: "Đặt vé thành công",
      });
      form.resetFields();
      setTicketPrice("0 VND");
      setTicketClasses([]);
      setSeats([]);
      setSelectedClass(null);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: error.message || "Có lỗi xảy ra khi đặt vé",
      });
    }
    console.log("Ticket data:", ticketData);
  };
  return (
    <Card title="Vé Chuyến Bay" style={{ width: 500, margin: "auto" }}>
      <Form form={form} layout="vertical" onFinish={handleSubmit}>
        <Form.Item label="Chuyến bay" name="flightCode">
          <Input placeholder="Chọn chuyến bay" disabled />
        </Form.Item>

        <Form.Item
          label="Loại vé"
          name="ticketType"
          rules={[{ required: true, message: "Vui lòng chọn loại vé" }]}
        >
          <Select
            placeholder="Chọn loại vé"
            disabled={!selectedFlight}
          >
            {ticketTypes.map((cls) => (
              <Option key={cls} value={cls}>
                {cls}
              </Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          label="Hành khách"
          name="passenger"
          rules={[{ required: true, message: "Vui lòng nhập tên hành khách" }]}
        >
          <Input placeholder="Nhập tên hành khách" />
        </Form.Item>

        <Form.Item
          label="CMND"
          name="idCard"
          rules={[{ required: true, message: "Vui lòng nhập CMND" }]}
        >
          <Input placeholder="Nhập CMND" />
        </Form.Item>

        <Form.Item
          label="Điện thoại"
          name="phone"
          rules={[{ required: true, message: "Vui lòng nhập số điện thoại" }]}
        >
          <Input placeholder="Nhập số điện thoại" />
        </Form.Item>

        <Form.Item label="Email" name="email">
          <Input placeholder="Nhập email (tùy chọn)" />
        </Form.Item>

        <Form.Item
          label="Hạng vé"
          name="ticketClass"
          rules={[{ required: true, message: "Vui lòng chọn hạng vé" }]}
        >
          <Select
            placeholder="Chọn hạng vé"
            onChange={handleClassChange}
            disabled={!selectedFlight}
          >
            {ticketClasses.map((cls) => (
              <Option key={cls} value={cls} defaultSeat={cls}>
                {cls}
              </Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          label="Số ghế"
          name="seat_number"
          rules={[{ required: true, message: "Vui lòng chọn ghế" }]}
        >
          <Select
            placeholder="Chọn số ghế"
            disabled={!selectedFlight || !selectedClass}
            onChange={handleSeatChange}
          >
            {seats
              .filter((seat) => seat.class_name === selectedClass)
              .map((seat) => (
                <Option
                  key={seat.seat_number}
                  value={seat.seat_number}
                  disabled={seat.is_booked}
                >
                  {seat.seat_number} {seat.is_booked ? "(Đã đặt)" : ""}
                </Option>
              ))}
          </Select>
        </Form.Item>

        <Form.Item label="Giá tiền">
          <Input value={ticketPrice} disabled />
        </Form.Item>

        <Button
          type="primary"
          className="w-full"
          htmlType="submit"
          disabled={!selectedFlight}
        >
          Bán vé
        </Button>
      </Form>
    </Card>
  );
};

export default TicketBooking;
