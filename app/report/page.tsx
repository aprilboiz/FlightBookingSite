"use client";

import { useState } from "react";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { toast } from "sonner";
import * as XLSX from "xlsx";
import { saveAs } from "file-saver";
import { getReportMonth, getReportYear } from "@/services/reportService";

interface FlightReport {
  flightCode: string;
  tickets: number;
  revenue: number;
}

interface MonthReport {
  month: number;
  flightCount: number;
  revenue: number;
}

interface ReportData {
  month?: number;
  year: number;
  flights?: FlightReport[];
  months?: MonthReport[];
  totalRevenue: number;
  totalTickets?: number;
  totalFlights?: number;
}

interface YearSheetData {
  Tháng: number | string;
  "Số chuyến bay"?: number;
  "Doanh thu (VNĐ)": number;
}

interface MonthSheetData {
  "Mã chuyến bay": string;
  "Số vé": number;
  "Doanh thu (VNĐ)": number;
}

export default function ReportPage() {
  const [month, setMonth] = useState<number>();
  const [year, setYear] = useState<number>();
  const [reportData, setReportData] = useState<ReportData | null>(null);
  const [loading, setLoading] = useState(false);

  const handleFetchReport = async () => {
    if (!month && !year) {
      toast.warning("Vui lòng nhập tháng hoặc năm hợp lệ!");
      return;
    }

    setLoading(true);
    try {
      let data;
      if (month) {
        data = await getReportMonth(month, year!);
      } else if (year) {
        data = await getReportYear(year);
      }
      setReportData(data);
    } catch (error) {
      toast.error("Lỗi khi lấy dữ liệu báo cáo!");
    } finally {
      setLoading(false);
    }
  };

  const exportToExcel = () => {
    if (!reportData) return;

    let sheetData: YearSheetData[] | MonthSheetData[];
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
        "Số chuyến bay": reportData.totalFlights || 0,
      });
    } else {
      // Báo cáo theo tháng
      sheetData = reportData.flights!.map((flight) => ({
        "Mã chuyến bay": flight.flightCode,
        "Số vé": flight.tickets,
        "Doanh thu (VNĐ)": flight.revenue,
      }));

      sheetData.push({
        "Mã chuyến bay": "",
        "Số vé": 0,
        "Doanh thu (VNĐ)": 0,
      });
      sheetData.push({
        "Mã chuyến bay": "TỔNG DOANH THU",
        "Số vé": 0,
        "Doanh thu (VNĐ)": reportData.totalRevenue,
      });
      sheetData.push({
        "Mã chuyến bay": "TỔNG VÉ BÁN",
        "Số vé": reportData.totalTickets || 0,
        "Doanh thu (VNĐ)": 0,
      });
    }

    const ws = XLSX.utils.json_to_sheet(sheetData);
    const wb = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(wb, ws, "Báo cáo thu nhập");

    const fileName = `BaoCao_ThuNhap_${month || year}.xlsx`;
    const wbout = XLSX.write(wb, { bookType: "xlsx", type: "array" });
    saveAs(new Blob([wbout], { type: "application/octet-stream" }), fileName);
  };

  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "calc(var(--spacing) * 72)",
          "--header-height": "calc(var(--spacing) * 12)",
        } as React.CSSProperties
      }
    >
      <AppSidebar variant="inset" />
      <SidebarInset>
        <SiteHeader />
        <div className="flex flex-1 flex-col">
          <div className="@container/main flex flex-1 flex-col gap-2">
            <div className="flex flex-col gap-4 py-4 md:gap-6 md:py-6">
              <div className="px-4 lg:px-6">
                <div className="w-full flex flex-col gap-5 items-center">
                  <h2 className="text-2xl font-bold">Báo cáo thu nhập</h2>

                  <div className="mb-4 text-center text-muted-foreground">
                    <i>
                      Hướng dẫn sử dụng: Nếu muốn thực hiện báo cáo tháng, hãy nhập đầy đủ{" "}
                      <strong>Tháng</strong> và <strong>Năm</strong>; nếu muốn thực hiện báo
                      cáo năm vui lòng chỉ nhập <strong>Năm</strong>
                    </i>
                  </div>

                  <div className="flex gap-4 items-center">
                    <div className="flex items-center gap-2">
                      <label htmlFor="month">Tháng:</label>
                      <Input
                        id="month"
                        type="number"
                        min={1}
                        max={12}
                        value={month}
                        onChange={(e) => setMonth(Number(e.target.value))}
                        placeholder="VD: 5"
                        className="w-24"
                      />
                    </div>
                    <div className="flex items-center gap-2">
                      <label htmlFor="year">Năm:</label>
                      <Input
                        id="year"
                        type="number"
                        min={2000}
                        value={year}
                        onChange={(e) => setYear(Number(e.target.value))}
                        placeholder="VD: 2025"
                        className="w-24"
                      />
                    </div>
                    <Button onClick={handleFetchReport}>Lấy báo cáo</Button>
                    {reportData && (
                      <Button variant="outline" onClick={exportToExcel}>
                        Xuất Excel
                      </Button>
                    )}
                  </div>

                  {loading ? (
                    <div className="flex justify-center">
                      <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
                    </div>
                  ) : reportData ? (
                    <div className="w-full">
                      {reportData.months ? (
                        // Báo cáo theo năm
                        <>
                          <h3 className="text-xl font-semibold mb-4">
                            Báo cáo năm {reportData.year}
                          </h3>
                          <div className="space-y-2 mb-4">
                            <p className="font-medium">
                              Tổng doanh thu: {reportData.totalRevenue} VNĐ
                            </p>
                            <p className="font-medium">
                              Tổng chuyến bay: {reportData.totalFlights}
                            </p>
                          </div>

                          <Table>
                            <TableHeader>
                              <TableRow>
                                <TableHead>Tháng</TableHead>
                                <TableHead>Số chuyến bay</TableHead>
                                <TableHead>Doanh thu (VNĐ)</TableHead>
                              </TableRow>
                            </TableHeader>
                            <TableBody>
                              {reportData.months.map((item) => (
                                <TableRow key={item.month}>
                                  <TableCell>{item.month}</TableCell>
                                  <TableCell>{item.flightCount}</TableCell>
                                  <TableCell>{item.revenue}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </>
                      ) : (
                        // Báo cáo theo tháng
                        <>
                          <h3 className="text-xl font-semibold mb-4">
                            Báo cáo tháng {reportData.month}
                          </h3>
                          <div className="space-y-2 mb-4">
                            <p className="font-medium">
                              Tổng doanh thu: {reportData.totalRevenue} VNĐ
                            </p>
                            <p className="font-medium">
                              Tổng số vé: {reportData.totalTickets}
                            </p>
                          </div>

                          <Table>
                            <TableHeader>
                              <TableRow>
                                <TableHead>Mã chuyến bay</TableHead>
                                <TableHead>Số vé</TableHead>
                                <TableHead>Doanh thu (VNĐ)</TableHead>
                              </TableRow>
                            </TableHeader>
                            <TableBody>
                              {reportData.flights?.map((flight) => (
                                <TableRow key={flight.flightCode}>
                                  <TableCell>{flight.flightCode}</TableCell>
                                  <TableCell>{flight.tickets}</TableCell>
                                  <TableCell>{flight.revenue}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </>
                      )}
                    </div>
                  ) : (
                    <p className="text-muted-foreground">Không có dữ liệu.</p>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
} 