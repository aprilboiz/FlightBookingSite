import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom';
import { FaPlane, FaPlaneDeparture, FaPlaneArrival  } from "react-icons/fa";
import { IoCashOutline, IoTime  } from "react-icons/io5";
import { RiPlaneLine } from "react-icons/ri";
import { CiStickyNote, CiCalendarDate  } from "react-icons/ci";
import { getFlightByCode } from '../services/flightService';
import { notification } from 'antd';
import { getPlaneByCode } from '../services/planeService';

const FlightDetail = () => {

  const { code } = useParams();
  const [flight, setFlight] = useState(null);
  const [plane, setPlane] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchDataFlightDetail();
  },[])

  const fetchDataFlightDetail = async () => {
    try {
      const data = await getFlightByCode(code);
      setFlight(data);
      fetchDataPlaneDetail(data.plane_code);
    } catch (error) {
      notification.error({
        message: 'Lỗi',
        description: error.message,
      });
    } finally {
      setLoading(false);
    }
  }
  
  const fetchDataPlaneDetail = async (planeCode) => {
    try {
      const data = await getPlaneByCode(planeCode);
      setPlane(data);
    } catch (error) {
      notification.error({
        message: 'Lỗi',
        description: error.message,
      });
    }  
  }

  if (loading || !plane || !flight) {
    return <div className="text-center mt-10">Đang tải dữ liệu...</div>;
  }

  const departureDate = new Date(flight.departure_date_time)
  const arrivalDate = new Date(departureDate.getTime() + flight.duration * 60000);
  const formatTime = (date) => {
    return date.toLocaleTimeString('vi-VN', { hour: '2-digit', minute: '2-digit' });
  }

  return (
    <div className='w-full flex flex-col items-center justify-center px-64'>
      <div className='flex items-center justify-between w-full m-10 text-xl'>
        <div>
          <h3 className='text-2xl font-medium'>{formatTime(departureDate)}</h3>
          <p className='font-light'>{flight.departure_airport}</p>
        </div>
        <div className='flex justify-center items-center flex-col font-light'>
          <p>{flight.duration} phút</p>
          <div className='flex items-center justify-center gap-2'>
            <div className='w-52 h-[1px] bg-black'></div>
            <FaPlane />
          </div>
          <p>{flight.intermediate_stop.length > 0 ? `${flight.intermediate_stop.length} điểm dừng` : `Bay thẳng`}</p>
        </div>
        <div>
          <h3 className='text-2xl font-medium'>{formatTime(arrivalDate)}</h3>
          <p className='font-light'>{flight.arrival_airport}</p>
        </div>
      </div>
      <hr className='bg-black w-3/4 h-[1px]'/>
      <h2 className='text-xl mt-10 font-light'>Thông tin chi tiết chuyến bay</h2>
      <div className='w-full flex items-start justify-between m-10'>
        <div className='flex flex-col items-start justify-center gap-10'>
          <div className='flex items-center justify-center gap-6'>
            <FaPlaneDeparture  className='text-3xl'/>
            <p className='text-base font-light'>{flight.departure_airport}</p>
          </div>
          <div className='flex items-center justify-center gap-6'>
            <FaPlaneArrival  className='text-3xl'/>
            <p className='text-base font-light'>{flight.arrival_airport}</p>
          </div>
          <div className='flex items-center justify-center gap-6'>
            <IoCashOutline  className='text-3xl'/>
            <p className='text-base font-light'>{flight.base_price} vnđ</p>
          </div>
          <div className='flex items-center justify-center gap-6'>
            <IoTime className='text-3xl'/>
            <p className='text-base font-light'>{flight.duration} phút</p>
          </div>
          <div className='flex items-center justify-center gap-6'>
            <CiCalendarDate className='text-3xl'/>
            <p className='text-base font-light'>{`${departureDate.getDate()} - ${departureDate.getMonth()} - ${departureDate.getFullYear()}`}</p>
          </div>
        </div>
        <div className='flex flex-col items-center justify-center gap-10'>
          <span className='flex items-center justify-center gap-2 text-base'>
            <p className='font-medium'>Mã máy bay:</p>
            <p className='font-light'>{plane.plane_code}</p>
          </span>
          <span className='flex items-center justify-center gap-2 text-base'>
            <p className='font-medium'>Tên máy bay bay:</p>
            <p className='font-light'>{plane.plane_name}</p>
          </span>
        </div>
      </div>
      <hr className='bg-black w-3/4 h-[1px]'/>
      {flight.intermediate_stop.length > 0 && (
        <>
          <h2 className='text-xl mt-10 font-light'>Thông tin chi tiết chuyến bay trung gian</h2>
          {flight.intermediate_stop.map((item, index) => (
            <div key={index} className='flex items-center justify-between w-full mt-10'>
              <div className='flex items-center justify-center gap-6'>
                <RiPlaneLine className='text-3xl'/>
                <p className='text-base font-light'>{item.stop_airport || "Không rõ"}</p>
              </div>
              <div className='flex items-center justify-center gap-6'>
                <IoTime className='text-3xl'/>
                <p className='text-base font-light'>{item.stop_duration || "Không rõ"} phút</p>
              </div>
              <div className='flex items-center justify-center gap-6'>
                <CiStickyNote className='text-3xl'/>
                <p className='text-base font-light'>{item.note || "Không có ghi chú"}</p>
              </div>
            </div>
          ))}
        </>
      )}
    </div>
  )
}

export default FlightDetail