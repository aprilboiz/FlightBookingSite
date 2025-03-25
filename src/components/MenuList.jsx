import React from 'react'
import { Menu } from 'antd'
import { HomeOutlined, FormOutlined, CreditCardOutlined } from '@ant-design/icons'
import { Link } from 'react-router-dom'

const MenuList = () => {
  return (
    <Menu theme="dark" className='h-[88vh] mt-[2rem] flex flex-col gap-[15px] text-[1rem] relative min-w-[220px]'>
        <Menu.Item key="home" icon={<HomeOutlined />}>
            <Link to="/">Trang chủ</Link>
        </Menu.Item>
        <Menu.Item key="calender-plane" icon={<FormOutlined />}>
            <Link to="/calender-plane">Lịch chuyến bay</Link>
        </Menu.Item>
        <Menu.Item key="list-plane" icon={<FormOutlined />}>
            <Link to="/list-plane">Danh sách chuyến bay</Link>
        </Menu.Item>
        <Menu.Item key="sale-ticket" icon={<CreditCardOutlined />}>
            <Link to="/sale-ticket">Vé máy bay</Link>
        </Menu.Item>
    </Menu>
  )
}

export default MenuList