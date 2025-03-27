import React from 'react'
import { Link } from 'react-router-dom'
const SideBar = ({ items }) => {

  return (
    <div className="w-64 min-h-screen p-4 border-r bg-gray-100 fixed">
      <h2 className="text-xl font-bold mb-4">TÀI LIỆU</h2>
      <ul className='pl-5'>
        {items.map((item, index) => (
          <li key={index} className="mb-2">
            <Link to={item.path} className=" hover:underline">
              {item.name}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default SideBar