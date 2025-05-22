"use client"

import * as React from "react"
import { Area, AreaChart, CartesianGrid, XAxis, YAxis, ResponsiveContainer } from "recharts"
import { getReportYear, getReportMonth } from "@/services/reportService"
import { toast } from "sonner"
import { FlightReport, MonthlyReportResponse, YearlyReportResponse } from "@/types/report"

import { useIsMobile } from "@/hooks/use-mobile"
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import {
  ToggleGroup,
  ToggleGroupItem,
} from "@/components/ui/toggle-group"

export const description = "Biểu đồ doanh thu theo tháng"

interface ChartDataItem {
  date: string
  revenue: number
  flightCount: number
  flightDetails?: FlightReport[]
}

export function ChartAreaInteractive() {
  const isMobile = useIsMobile()
  const [timeRange, setTimeRange] = React.useState("12m")
  const [chartData, setChartData] = React.useState<ChartDataItem[]>([])
  const currentYear = new Date().getFullYear()
  const currentMonth = new Date().getMonth() + 1

  React.useEffect(() => {
    if (isMobile) {
      setTimeRange("3m")
    }
  }, [isMobile])

  React.useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch both yearly and current month data
        const [yearlyResponse, monthlyResponse] = await Promise.all([
          getReportYear(currentYear),
          getReportMonth(currentMonth, currentYear)
        ]) as [YearlyReportResponse, MonthlyReportResponse]

        // Format yearly data
        let formattedYearlyData: ChartDataItem[] = yearlyResponse.months.map(month => ({
          date: `${currentYear}-${String(month.month).padStart(2, '0')}-01`,
          revenue: month.revenue,
          flightCount: month.flightCount
        }))

        // Add current month's detailed data if available
        if (monthlyResponse.flights && monthlyResponse.flights.length > 0) {
          const monthData = formattedYearlyData.find(
            item => new Date(item.date).getMonth() + 1 === currentMonth
          )
          if (monthData) {
            monthData.flightDetails = monthlyResponse.flights
          }
        }

        setChartData(formattedYearlyData)
      } catch (error) {
        toast.error("Không thể tải dữ liệu doanh thu")
      }
    }

    fetchData()
  }, [currentYear, currentMonth])

  const chartConfig = {
    revenue: {
      label: "Doanh thu",
      color: "var(--primary)",
    }
  } satisfies ChartConfig

  const filteredData = chartData.filter((item) => {
    const date = new Date(item.date)
    const referenceDate = new Date(`${currentYear}-12-31`)
    let monthsToSubtract = 12
    if (timeRange === "6m") {
      monthsToSubtract = 6
    } else if (timeRange === "3m") {
      monthsToSubtract = 3
    }
    const startDate = new Date(referenceDate)
    startDate.setMonth(startDate.getMonth() - monthsToSubtract)
    return date >= startDate
  })

  return (
    <Card className="@container/card">
      <CardHeader>
        <CardTitle>Doanh thu theo tháng</CardTitle>
        <CardDescription>
          <span className="hidden @[540px]/card:block">
            Doanh thu theo tháng trong năm {currentYear}
          </span>
          <span className="@[540px]/card:hidden">Năm {currentYear}</span>
        </CardDescription>
        <CardAction>
          <ToggleGroup
            type="single"
            value={timeRange}
            onValueChange={setTimeRange}
            variant="outline"
            className="hidden *:data-[slot=toggle-group-item]:!px-4 @[767px]/card:flex"
          >
            <ToggleGroupItem value="12m">12 tháng</ToggleGroupItem>
            <ToggleGroupItem value="6m">6 tháng</ToggleGroupItem>
            <ToggleGroupItem value="3m">3 tháng</ToggleGroupItem>
          </ToggleGroup>
          <Select value={timeRange} onValueChange={setTimeRange}>
            <SelectTrigger
              className="flex w-40 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate @[767px]/card:hidden"
              size="sm"
              aria-label="Select a value"
            >
              <SelectValue placeholder="12 tháng" />
            </SelectTrigger>
            <SelectContent className="rounded-xl">
              <SelectItem value="12m" className="rounded-lg">
                12 tháng
              </SelectItem>
              <SelectItem value="6m" className="rounded-lg">
                6 tháng
              </SelectItem>
              <SelectItem value="3m" className="rounded-lg">
                3 tháng
              </SelectItem>
            </SelectContent>
          </Select>
        </CardAction>
      </CardHeader>
      <CardContent className="px-2 pt-4 sm:px-6 sm:pt-6">
        <div className="h-[300px] w-full">
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart data={filteredData}>
              <defs>
                <linearGradient id="fillRevenue" x1="0" y1="0" x2="0" y2="1">
                  <stop
                    offset="5%"
                    stopColor="hsl(var(--primary))"
                    stopOpacity={0.8}
                  />
                  <stop
                    offset="95%"
                    stopColor="hsl(var(--primary))"
                    stopOpacity={0.1}
                  />
                </linearGradient>
              </defs>
              <CartesianGrid strokeDasharray="3 3" vertical={false} />
              <XAxis
                dataKey="date"
                tickLine={false}
                axisLine={false}
                tickMargin={8}
                minTickGap={32}
                tickFormatter={(value) => {
                  const date = new Date(value)
                  return date.toLocaleDateString("vi-VN", {
                    month: "short",
                    year: "numeric"
                  })
                }}
              />
              <YAxis
                tickLine={false}
                axisLine={false}
                tickMargin={8}
                tickFormatter={(value) => {
                  return `${(value / 1000000).toFixed(1)}M`
                }}
              />
              <Area
                dataKey="revenue"
                type="monotone"
                fill="url(#fillRevenue)"
                stroke="hsl(var(--primary))"
                strokeWidth={2}
              />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  )
}
