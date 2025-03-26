import React from 'react'
import {
    Input,
    Table,
} from 'antd'
import { data } from 'react-router-dom'
import TicketBooking from './TicketBooking'

const { Search } = Input

const SaleTicket = () => {

  const columns = [
    { title: 'Mã chuyến bay', dataIndex: 'code', key: 'code' },
    { title: "Sân bay từ", dataIndex: "from", key: "from" },
    { title: "Sân bay đến", dataIndex: "to", key: "to" },
    { title: "Sân bay trung gian", dataIndex: "transit", key: "transit" },
    { title: "Ngày - giờ khởi hành", dataIndex: "date", key: "date" },
    { title: "Giá vé", dataIndex: "price", key: "price" },
    { title: "Số lượng ghế", dataIndex: "quantity", key: "quantity" },
  ]

  const data = [
    {  key: '1', code: 'RUA001', from: 'Hà Nội', to: 'TP.Hồ Chí Minh', transit: 'Đà Nẵng', date: '01/01/2022 08:00', price: '1.000.000', quantity: '100' },
    {  key: '2', code: 'RUA002', from: 'TP.Hồ Chí Minh', to: 'Hà Nội', transit: 'Đà Nẵng', date: '01/01/2022 08:00', price: '1.000.000', quantity: '100' },
    {  key: '3', code: 'RUA003', from: 'Hà Nội', to: 'Đà Nẵng', transit: '', date: '01/01/2022 08:00', price: '1.000.000', quantity: '100' },
  ]

  return (
    <div className='flex justify-between items-start gap-5'>
        <div className='w-3/4'>
            <div><Search placeholder='Tìm kiếm mã chuyến bay' enterButton/></div>
            <div>
                <Table
                    columns={columns}
                    dataSource={data}
                    pagination={false}
                    bordered
                    scroll={{y:300}}
                    className="mt-5"
                />
            </div>
        </div>

        <div>
            <TicketBooking />
        </div>
    </div>
  )
}

export default SaleTicket