import React from 'react'
import { Layout, List } from 'antd'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Logo from './components/Logo'
import MenuList from './components/MenuList'
import CalenderPlane from './components/CalenderPlane'
import ListPlane from './components/ListPlane'
import SaleTicket from './components/SaleTicket'

const { Header, Sider, Content } = Layout
function App() {

  return (
    <Router>
      <Layout>
        <Sider className='text-white'>
          <Logo />
          <MenuList />
        </Sider>
        <Layout>
          <Header className='bg-white'>Header</Header>
          <Content className='p-10'>
            <Routes>
              <Route path="/calender-plane" element={<CalenderPlane />} />
              <Route path="/list-plane" element={<ListPlane />} />
              <Route path="/sale-ticket" element={<SaleTicket />} />
            </Routes>
          </Content>
        </Layout >
      </Layout>
    </Router>
  )
}

export default App
