import React, { useEffect } from "react";
import { Form, InputNumber, Button, Card, notification } from "antd";

import { getParameter, updateParameter } from "../services/parameterService";

const RegulationForm = () => {
  const [form] = Form.useForm();

  useEffect(() => {
    fetchRegulation();
  }, []);

  const fetchRegulation = async () => {
    try {
      const data = await getParameter();
      form.setFieldsValue(data);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải quy định.",
      });
    }
  };

  const onFinish = async (values) => {
    try {
      await updateParameter(values);
      notification.success({
        message: "Thành công",
        description: "Cập nhật quy định thành công!",
      });
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể cập nhật quy định.",
      });
    }
  };

  return (
    <Card title="Cập nhật Quy định chuyến bay" style={{ maxWidth: 600, margin: "auto" }}>
      <Form form={form} layout="vertical" onFinish={onFinish}>
        <Form.Item label="Số lượng sân bay" name="number_of_airports" rules={[{ required: true }]}>
          <InputNumber min={1} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="Thời gian bay tối thiểu (phút)" name="min_flight_duration" rules={[{ required: true }]}>
          <InputNumber min={1} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="Số điểm dừng tối đa" name="max_intermediate_stops" rules={[{ required: true }]}>
          <InputNumber min={0} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="TG dừng tối thiểu (phút)" name="min_intermediate_stop_duration" rules={[{ required: true }]}>
          <InputNumber min={1} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="TG dừng tối đa (phút)" name="max_intermediate_stop_duration" rules={[{ required: true }]}>
          <InputNumber min={1} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="Số hạng vé tối đa" name="max_ticket_classes" rules={[{ required: true }]}>
          <InputNumber min={1} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="TG mua vé chậm nhất (ngày)" name="latest_ticket_purchase_time" rules={[{ required: true }]}>
          <InputNumber min={0} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item label="TG hủy vé tối thiểu (phút trước giờ bay)" name="ticket_cancellation_time" rules={[{ required: true }]}>
          <InputNumber min={0} style={{ width: "100%" }} />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit" block>
            Cập nhật quy định
          </Button>
        </Form.Item>
      </Form>
    </Card>
  );
};

export default RegulationForm;
