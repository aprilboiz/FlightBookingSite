import React from 'react'
import { Layout } from 'antd'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Logo from './components/Logo'
import MenuList from './components/MenuList'
import CalenderPlane from './components/CalenderPlane'

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
            </Routes>
          </Content>
        </Layout >
      </Layout>
    </Router>
  )
}

export default App
