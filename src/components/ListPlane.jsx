import React, { useState, useEffect } from "react";
import { getFlights } from "../services/flightService";
// import { getAirports } from "../services/airportService.js";
import { useNavigate } from "react-router-dom";
import {
  Table,
  Button,
  // Modal,
  Form,
  // Input,
  // InputNumber,
  // DatePicker,
  // Space,
  notification,
  // Select,
} from "antd";
import {
  // EditOutlined,
  // DeleteOutlined,
  // PlusOutlined,
  // MinusCircleOutlined,
  InfoCircleOutlined,
} from "@ant-design/icons";
import dayjs from "dayjs";

const ListPlane = () => {
  const [flights, setFlights] = useState([]);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [form] = Form.useForm();
  const [editingFlight, setEditingFlight] = useState(null);

  const [airports, setAirports] = useState([]);

  const navigate = useNavigate();

  useEffect(() => {
    fetchFlights();
    // fetchAirports();
  }, []);

  const fetchFlights = async () => {
    try {
      const data = await getFlights();
      const formatted = data.map((item, index) => ({
        key: index,
        flightCode: item.flight_code,
        departureAirport: item.departure_airport,
        arrivalAirport: item.arrival_airport,
        departureTime: item.departure_date,
        price: item.base_price,
        intermediateStops: item.intermediate_stops || [],
      }));
      setFlights(formatted);
    } catch (error) {
      notification.error({
        message: "Lỗi",
        description: "Không thể tải danh sách chuyến bay",
      });
    }
  };

  // const fetchAirports = async () => {
  //   try {
  //     const data = await getAirports();
  //     setAirports(data);
  //   } catch (err) {
  //     notification.error({
  //       message: "Lỗi",
  //       description: "Không thể lấy danh sách sân bay",
  //     });
  //   }
  // };

  // const handleEdit = (record) => {
  //   setEditingFlight(record);
  //   form.setFieldsValue({
  //     ...record,
  //     departureTime: dayjs(record.departureTime),
  //   });
  //   setIsModalVisible(true);
  // };

  // const handleDelete = (record) => {
  //   console.log("Xoá:", record);
  //   // Thêm logic xoá ở đây
  // };

  // const handleOk = () => {
  //   form.validateFields().then((values) => {
  //     console.log("Giá trị sau khi chỉnh sửa:", values);
  //     // Gửi API cập nhật ở đây nếu cần
  //     setIsModalVisible(false);
  //   });
  // };

  const columns = [
    {
      title: "Mã chuyến bay",
      dataIndex: "flightCode",
      key: "flightCode",
    },
    {
      title: "Sân bay đi",
      dataIndex: "departureAirport",
      key: "departureAirport",
    },
    {
      title: "Sân bay đến",
      dataIndex: "arrivalAirport",
      key: "arrivalAirport",
    },
    {
      title: "Thời gian khởi hành",
      dataIndex: "departureTime",
      key: "departureTime",
      render: (text) => dayjs(text).format("YYYY-MM-DD HH:mm:ss"),
    },
    {
      title: "Giá vé",
      dataIndex: "price",
      key: "price",
    },
    {
      title: "Hành động",
      key: "action",
      render: (_, record) => (
        // <span className="flex gap-2">
        //   <Button icon={<EditOutlined />} onClick={() => handleEdit(record)} />
        //   <Button
        //     icon={<DeleteOutlined />}
        //     danger
        //     onClick={() => handleDelete(record)}
        //   />
        // </span>
        <Button
          icon={<InfoCircleOutlined />}
          onClick={() => navigate(`/list-plane/${record.flightCode}`)}
        />
      ),
    },
  ];

  return (
    <>
      <div className="w-full p-5 flex flex-col gap-5 items-center">
        <h2>DANH SÁCH CHUYẾN BAY</h2>
        <div className="w-3/4">
          <Table dataSource={flights} columns={columns} />
        </div>
      </div>

      {/* <Modal
        title="Chỉnh sửa chuyến bay"
        open={isModalVisible}
        onOk={handleOk}
        onCancel={() => setIsModalVisible(false)}
        width={700}
        okText="Lưu"
        cancelText="Hủy"
      >
        <Form form={form} layout="vertical">
          <Form.Item
            name="departureAirport"
            label="Sân bay đi"
            rules={[{ required: true }]}
          >
            <Select placeholder="Chọn sân bay">
              {airports.map((airport) => (
                <Option key={airport.airport_code} value={airport.airport_code}>
                  {airport.airport_name}
                </Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item
            name="arrivalAirport"
            label="Sân bay đến"
            rules={[{ required: true, message: "Vui lòng chọn sân bay đến" }]}
          >
            <Select placeholder="Chọn sân bay">
              {airports.map((airport) => (
                <Option key={airport.airport_code} value={airport.airport_code}>
                  {airport.airport_name}
                </Option>
              ))}
            </Select>
          </Form.Item>
          <Form.Item
            name="departureTime"
            label="Thời gian khởi hành"
            rules={[{ required: true }]}
          >
            <DatePicker
              showTime
              format="YYYY-MM-DD HH:mm:ss"
              className="w-full"
            />
          </Form.Item>
          <Form.Item name="price" label="Giá vé" rules={[{ required: true }]}>
            <InputNumber min={0} className="w-full" />
          </Form.Item>

          <Form.List name="intermediateStops">
            {(fields, { add, remove }) => (
              <div>
                <label className="font-semibold">Sân bay trung gian</label>
                {fields.map(({ key, name, ...restField }) => (
                  <Space
                    key={key}
                    style={{ display: "flex", marginBottom: 8 }}
                    align="baseline"
                  >
                    <Select
                      placeholder="Chọn sân bay trung gian"
                      className="w-full"
                      value={stop.airport}
                      onChange={(value) => {
                        const updated = [...stopAirports];
                        updated[index].airport = value;
                        setStopAirports(updated);
                      }}
                    >
                      {airports.map((a) => (
                        <Option key={a.airport_code} value={a.airport_code}>
                          {a.airport_name}
                        </Option>
                      ))}
                    </Select>
                    <Form.Item
                      {...restField}
                      name={[name, "stop_duration"]}
                      rules={[
                        { required: true, message: "Nhập thời gian dừng" },
                      ]}
                    >
                      <InputNumber placeholder="Thời gian (phút)" />
                    </Form.Item>
                    <Form.Item {...restField} name={[name, "note"]}>
                      <Input placeholder="Ghi chú" />
                    </Form.Item>
                    <MinusCircleOutlined onClick={() => remove(name)} />
                  </Space>
                ))}
                <Button
                  type="dashed"
                  onClick={() => add()}
                  block
                  icon={<PlusOutlined />}
                >
                  Thêm sân bay trung gian
                </Button>
              </div>
            )}
          </Form.List>
        </Form>
      </Modal> */}
    </>
  );
};

export default ListPlane;
