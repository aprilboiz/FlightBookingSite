"use client"

import { useEffect, useState } from "react"
import { IconTrendingDown, IconTrendingUp } from "@tabler/icons-react"
import { getReportYear, getReportMonth } from "@/services/reportService"
import { toast } from "sonner"

import { Badge } from "@/components/ui/badge"
import {
  Card,
  CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"

export function SectionCards() {
  const [totalRevenue, setTotalRevenue] = useState<number>(0)
  const [monthlyRevenue, setMonthlyRevenue] = useState<number>(0)
  const currentYear = new Date().getFullYear()
  const currentMonth = new Date().getMonth() + 1 // getMonth() returns 0-11

  useEffect(() => {
    const fetchReports = async () => {
      try {
        const [yearlyResponse, monthlyResponse] = await Promise.all([
          getReportYear(currentYear),
          getReportMonth(currentMonth, currentYear)
        ])
        setTotalRevenue(yearlyResponse.totalRevenue)
        setMonthlyRevenue(monthlyResponse.totalRevenue)
      } catch (error) {
        toast.error("Không thể tải dữ liệu doanh thu")
      }
    }

    fetchReports()
  }, [currentYear, currentMonth])

  return (
    <div className="*:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card dark:*:data-[slot=card]:bg-card grid grid-cols-1 gap-4 px-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:shadow-xs lg:px-6 @xl/main:grid-cols-2 @5xl/main:grid-cols-4">
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Tổng doanh thu {currentYear}</CardDescription>
          <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
            {totalRevenue.toLocaleString('vi-VN')} VNĐ
          </CardTitle>
          <CardAction>
            <Badge variant="outline">
              <IconTrendingUp />
              {currentYear}
            </Badge>
          </CardAction>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium">
            Doanh thu năm {currentYear} <IconTrendingUp className="size-4" />
          </div>
          <div className="text-muted-foreground">
            Tổng doanh thu từ tất cả các chuyến bay
          </div>
        </CardFooter>
      </Card>
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Doanh thu tháng {currentMonth}/{currentYear}</CardDescription>
          <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
            {monthlyRevenue.toLocaleString('vi-VN')} VNĐ
          </CardTitle>
          <CardAction>
            <Badge variant="outline">
              <IconTrendingUp />
              Tháng {currentMonth}
            </Badge>
          </CardAction>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium">
            Doanh thu tháng {currentMonth} <IconTrendingUp className="size-4" />
          </div>
          <div className="text-muted-foreground">
            Tổng doanh thu từ các chuyến bay trong tháng
          </div>
        </CardFooter>
      </Card>   
    </div>
  )
}
