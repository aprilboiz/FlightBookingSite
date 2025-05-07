import React, { useState, useEffect } from "react";
import { Input, Table, notification } from "antd";
import TicketBooking from "./TicketBooking";
import { getFlights } from "../services/flightService.js";

const { Search } = Input;

const SaleTicket = () => {
  const [flights, setFlights] = useState([]);
  const [filteredFlights, setFilteredFlights] = useState([]);
  const [selectedFlight, setSelectedFlight] = useState(null);

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      const data = await getFlights();
      setFlights(data);
      setFilteredFlights(data);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải danh sách chuyến bay",
      });
    }
  };

  const handleSearch = (value) => {
    const result = flights.filter((flight) =>
      flight.flight_code.toLowerCase().includes(value.toLowerCase())
    );
    setFilteredFlights(result);
  };

  const handleRowClick = (record) => {
    setSelectedFlight(record);
  };

  return (
    <div className="flex justify-between items-start gap-5">
      <div className="w-3/4">
        <Input
          placeholder="Tìm kiếm mã chuyến bay"
          allowClear
          onChange={(e) => handleSearch(e.target.value)}
        />

        <Table
          dataSource={filteredFlights}
          rowKey={(record) => record.flight_code}
          pagination={false}
          onRow={(record) => ({
            onClick: () => handleRowClick(record),
          })}
          columns={[
            {
              title: "Mã chuyến bay",
              dataIndex: "flight_code",
              key: "flight_code",
            },
            {
              title: "Điểm đi",
              dataIndex: "departure_airport",
              key: "departure_airport",
            },
            {
              title: "Điểm đến",
              dataIndex: "arrival_airport",
              key: "arrival_airport",
            },
            {
              title: "Thời gian",
              dataIndex: "departure_date_time",
              key: "departure_date_time",
            },
          ]}
        />
      </div>
      <div>
        <TicketBooking selectedFlight={selectedFlight} />
      </div>
    </div>
  );
};

export default SaleTicket;