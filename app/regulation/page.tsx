"use client";

import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { toast } from "sonner";
import { getParameter, updateParameter, FlightParameter } from "@/services/parameterService";

const defaultValues: FlightParameter = {
  number_of_airports: 0,
  min_flight_duration: 0,
  max_intermediate_stops: 0,
  min_intermediate_stop_duration: 0,
  max_intermediate_stop_duration: 0,
  max_ticket_classes: 0,
  latest_ticket_purchase_time: 0,
  ticket_cancellation_time: 0
};

export default function RegulationPage() {
  const form = useForm<FlightParameter>({
    defaultValues
  });

  useEffect(() => {
    fetchRegulation();
  }, []);

  const fetchRegulation = async () => {
    try {
      const data = await getParameter();
      form.reset(data);
    } catch (error) {
      toast.error("Không thể tải quy định");
    }
  };

  const onSubmit = async (values: FlightParameter) => {
    try {
      await updateParameter(values);
      toast.success("Cập nhật quy định thành công!");
    } catch (error) {
      toast.error("Không thể cập nhật quy định");
    }
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
                  <Card className="w-full max-w-2xl">
                    <CardHeader>
                      <CardTitle>Cập nhật Quy định chuyến bay</CardTitle>
                    </CardHeader>
                    <CardContent>
                      <Form {...form}>
                        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
                          <FormField
                            control={form.control}
                            name="number_of_airports"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>Số lượng sân bay</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={1}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="min_flight_duration"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>Thời gian bay tối thiểu (phút)</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={1}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="max_intermediate_stops"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>Số điểm dừng tối đa</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={0}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="min_intermediate_stop_duration"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>TG dừng tối thiểu (phút)</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={1}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="max_intermediate_stop_duration"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>TG dừng tối đa (phút)</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={1}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="max_ticket_classes"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>Số hạng vé tối đa</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={1}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="latest_ticket_purchase_time"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>TG mua vé chậm nhất (ngày)</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={0}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <FormField
                            control={form.control}
                            name="ticket_cancellation_time"
                            render={({ field }) => (
                              <FormItem>
                                <FormLabel>TG hủy vé tối thiểu (phút trước giờ bay)</FormLabel>
                                <FormControl>
                                  <Input
                                    type="number"
                                    min={0}
                                    {...field}
                                    onChange={(e) => field.onChange(Number(e.target.value))}
                                  />
                                </FormControl>
                                <FormMessage />
                              </FormItem>
                            )}
                          />

                          <Button type="submit" className="w-full">
                            Cập nhật quy định
                          </Button>
                        </form>
                      </Form>
                    </CardContent>
                  </Card>
                </div>
              </div>
            </div>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
} 