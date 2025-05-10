import React, { useState } from "react";
import { getReportMonth, getReportYear } from "../services/reportService.js";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import {
  Form,
  InputNumber,
  Button,
  Table,
  Spin,
  Typography,
  Space,
  message,
} from "antd";

const { Title, Text } = Typography;

const ReportPage = () => {
  const [month, setMonth] = useState();
  const [year, setYear] = useState();
  const [reportData, setReportData] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleFetchReport = async () => {
    if (!month && !year) {
      message.warning("Vui lòng nhập tháng hoặc năm hợp lệ!");
      return;
    }

    setLoading(true);
    try {
      let data;
      if (month) {
        data = await getReportMonth(month, year);
      } else if (year) {
        data = await getReportYear(year);
      }
      setReportData(data);
    } catch (error) {
      message.error("Lỗi khi lấy dữ liệu báo cáo!");
    } finally {
      setLoading(false);
    }
  };

  const exportToExcel = () => {
    if (!reportData) return;

    let sheetData;
    if (reportData.months) {
      // Báo cáo theo năm
      sheetData = reportData.months.map((item) => ({
        Tháng: item.month,
        "Số chuyến bay": item.flightCount,
        "Doanh thu (VNĐ)": item.revenue,
      }));

      sheetData.push({
        Tháng: "TỔNG DOANH THU",
        "Doanh thu (VNĐ)": reportData.totalRevenue,
      });
      sheetData.push({
        Tháng: "TỔNG CHUYẾN BAY",
        "Số chuyến bay": reportData.totalFlights,
      });
    } else {
      // Báo cáo theo tháng
      sheetData = reportData.flights.map((flight) => ({
        "Mã chuyến bay": flight.flightCode,
        "Số vé": flight.tickets,
        "Doanh thu (VNĐ)": flight.revenue,
      }));

      sheetData.push({});
      sheetData.push({
        "Mã chuyến bay": "TỔNG DOANH THU",
        "Doanh thu (VNĐ)": reportData.totalRevenue,
      });
      sheetData.push({
        "Mã chuyến bay": "TỔNG VÉ BÁN",
        "Số vé": reportData.totalTickets,
      });
    }

    const ws = XLSX.utils.json_to_sheet(sheetData);
    const wb = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(wb, ws, "Báo cáo thu nhập");

    const fileName = `BaoCao_ThuNhap_${month || year}.xlsx`;
    const wbout = XLSX.write(wb, { bookType: "xlsx", type: "array" });
    saveAs(new Blob([wbout], { type: "application/octet-stream" }), fileName);
  };

  const columnsMonth = [
    { title: "Mã chuyến bay", dataIndex: "flightCode", key: "flightCode" },
    { title: "Số vé", dataIndex: "tickets", key: "tickets" },
    { title: "Doanh thu (VNĐ)", dataIndex: "revenue", key: "revenue" },
  ];

  const columnsYear = [
    { title: "Tháng", dataIndex: "month", key: "month" },
    { title: "Số chuyến bay", dataIndex: "flightCount", key: "flightCount" },
    { title: "Doanh thu (VNĐ)", dataIndex: "revenue", key: "revenue" },
  ];

  return (
    <div style={{ padding: 24 }}>
      <Title level={2} style={{ textAlign: "center" }}>
        Báo cáo thu nhập
      </Title>

      <div className="mb-4">
        <i>
          Hướng dẫn sử dụng: Nếu muốn thực hiện báo cáo tháng, hãy nhập đầy đủ{" "}
          <strong>Tháng</strong> và <strong>Năm</strong>; nếu muốn thực hiện báo
          cáo năm vui lòng chỉ nhập <strong>Năm</strong>
        </i>
      </div>

      <Form layout="inline" style={{ marginBottom: 24 }}>
        <Form.Item label="Tháng">
          <InputNumber
            min={1}
            max={12}
            value={month}
            onChange={setMonth}
            placeholder="VD: 5"
          />
        </Form.Item>
        <Form.Item label="Năm">
          <InputNumber
            min={2000}
            value={year}
            onChange={setYear}
            placeholder="VD: 2025"
          />
        </Form.Item>
        <Form.Item>
          <Button type="primary" onClick={handleFetchReport}>
            Lấy báo cáo
          </Button>
        </Form.Item>
        {reportData && (
          <Form.Item>
            <Button onClick={exportToExcel}>Xuất Excel</Button>
          </Form.Item>
        )}
      </Form>

      {loading ? (
        <Spin tip="Đang tải dữ liệu..." />
      ) : reportData ? (
        <>
          {reportData.months ? (
            // Báo cáo theo năm
            <>
              <Title level={4}>Báo cáo năm {reportData.year}</Title>
              <Space direction="vertical" style={{ marginBottom: 16 }}>
                <Text strong>
                  Tổng doanh thu: {reportData.totalRevenue} VNĐ
                </Text>
                <Text strong>Tổng chuyến bay: {reportData.totalFlights}</Text>
              </Space>

              <Table
                dataSource={reportData.months}
                columns={columnsYear}
                rowKey={(record) => record.month}
                pagination={false}
                bordered
              />
            </>
          ) : (
            // Báo cáo theo tháng
            <>
              <Title level={4}>Báo cáo tháng {reportData.month}</Title>
              <Space direction="vertical" style={{ marginBottom: 16 }}>
                <Text strong>
                  Tổng doanh thu: {reportData.totalRevenue} VNĐ
                </Text>
                <Text strong>Tổng số vé: {reportData.totalTickets}</Text>
              </Space>

              <Table
                dataSource={reportData.flights}
                columns={columnsMonth}
                rowKey={(record) => record.flightCode}
                pagination={false}
                bordered
              />
            </>
          )}
        </>
      ) : (
        <Text>Không có dữ liệu.</Text>
      )}
    </div>
  );
};

export default ReportPage;
