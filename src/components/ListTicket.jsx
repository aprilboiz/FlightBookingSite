import React, { useState, useEffect, use } from "react";
import { getTickets, updateStatusTicket } from "../services/ticketService.js";
import { getFlightByCode } from "../services/flightService.js";

import { notification, Table, Button } from "antd";

import dayjs from "dayjs";
const ListTicket = () => {
  const [tickets, setTickets] = useState([]);

  useEffect(() => {
    fetchTickets();
  }, []);

  const fetchTickets = async () => {
    try {
      const data = await getTickets();

      const formatted = await Promise.all(
        data.map(async (item, index) => {
          const flight = await getFlightByCode(item.flight_code);

          return {
            key: index,
            ticketCode: item.id,
            flightCode: item.flight_code,
            passengerName: item.full_name,
            phoneNumber: item.phone_number,
            email: item.email,
            idCard: item.id_card,
            ticketPrice: item.price + " VNĐ",
            paymentStatus: item.ticket_status,
            ticketType: item.booking_type,
            departureDate: flight.departure_date_time
              ? dayjs(flight.departure_date_time).format("YYYY-MM-DD HH:mm:ss")
              : "Không có dữ liệu",
          };
        })
      );

      setTickets(formatted);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải danh sách vé máy bay",
      });
    }
  };

  const handleUpdateStatus = async (ticketCode, departureDate) => {
    const currentDate = new Date();
    const flightDate = new Date(departureDate);
    console.log("Current Date:", currentDate);
    console.log("Flight Date:", flightDate);
    if (flightDate <= currentDate) {
      try {
        await updateStatusTicket(ticketCode, "EXPIRED");
        notification.success({
          message: "Cập nhật trạng thái thành công",
          description: "Trạng thái vé đã được cập nhật",
        });
        fetchTickets();
      } catch (error) {
        notification.error({
          message: "Lỗi",
          description: "Không thể cập nhật trạng thái vé",
        });
      }
    }
  };

  const handleCancelTicket = async (ticketCode) => {
    try {
      await updateStatusTicket(ticketCode, "CANCELLED");
      notification.success({
        message: "Hủy vé thành công",
        description: "Trạng thái vé đã được cập nhật",
      });
      fetchTickets();
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể hủy vé",
      });
    }
  };

  const columns = [
    {
      title: "Mã vé",
      dataIndex: "ticketCode",
      key: "ticketCode",
    },
    {
      title: "Mã chuyến bay",
      dataIndex: "flightCode",
      key: "flightCode",
    },
    {
      title: "Tên hành khách",
      dataIndex: "passengerName",
      key: "passengerName",
    },
    {
      title: "Số điện thoại",
      dataIndex: "phoneNumber",
      key: "phoneNumber",
    },
    {
      title: "Email",
      dataIndex: "email",
      key: "email",
    },
    {
      title: "CMND/CCCD",
      dataIndex: "idCard",
      key: "idCard",
    },
    {
      title: "Giá vé",
      dataIndex: "ticketPrice",
      key: "ticketPrice",
    },
    {
      title: "Trạng thái vé",
      dataIndex: "paymentStatus",
      key: "paymentStatus",
    },
    {
      title: "Ngày khởi hành",
      dataIndex: "departureDate",
      key: "departureDate",
    },
    {
      title: "Loại vé",
      dataIndex: "ticketType",
      key: "ticketType",
        render: (text) => {
        if (text === "TICKET") return "Vé mua trực tiếp";
        if (text === "PLACE_ORDER") return "Vé đặt trước";
        return "Không xác định";
        }
    },
    {
      title: "Thao tác",
      key: "actions",
      render: (_, record) => (
        <div>
          <Button
            onClick={() =>
              handleUpdateStatus(record.ticketCode, record.departureDate)
            }
            type="primary"
            style={{ marginRight: 10 }}
            disabled={record.paymentStatus === "CANCELLED"}
          >
            Cập nhật trạng thái
          </Button>
          <Button
            onClick={() => handleCancelTicket(record.ticketCode)}
            danger
            disabled={record.paymentStatus === "CANCELLED"}
          >
            Hủy vé
          </Button>
        </div>
      ),
    },
  ];

  return (
    <>
      <div className="w-full p-5 flex flex-col gap-5 items-center">
        <h2>DANH SÁCH VÉ MÁY BAY</h2>
        <div className="w-3/4">
          <Table dataSource={tickets} columns={columns} pagination={false} />
        </div>
      </div>
    </>
  );
};

export default ListTicket;
