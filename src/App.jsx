import React from "react";
import { Layout, List } from "antd";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Logo from "./components/Logo";
import MenuList from "./components/MenuList";
import CalenderPlane from "./components/CalenderPlane";
import ListPlane from "./components/ListPlane";
import SaleTicket from "./components/SaleTicket";
import Test from "./components/test";
import FlightDetail from "./components/FlightDetail";
import ListTicket from "./components/ListTicket";
import RegulationForm from "./components/RegulationForm";
import ReportPage from "./components/ReportPage";
import Welcome from "./components/Welcome";

const { Header, Sider, Content } = Layout;
function App() {
  return (
    <Router>
      <div className="flex justify-center items-center">
        <div className="w-[300px] h-screen border-l-2">
          <Logo />
          <hr />
          <MenuList />
        </div>
        <div className="w-full h-screen">
          <Routes>
            <Route path="/calender-plane" element={<CalenderPlane />} />
            <Route path="/list-plane" element={<ListPlane />} />
            <Route path="/list-plane/:code" element={<FlightDetail />} />
            <Route path="/sale-ticket" element={<SaleTicket />} />
            <Route path="/list-ticket" element={<ListTicket />} />
            <Route path="/regulation" element={<RegulationForm />} />
            <Route path="/report" element={<ReportPage />} />
            <Route path="/" element={<Welcome />} />
          </Routes>
        </div>
        <div className="fixed bottom-4 right-4 z-50">
          <Test />
        </div>
      </div>
    </Router>
  );
}

export default App;
