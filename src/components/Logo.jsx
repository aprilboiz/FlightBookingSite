import React from 'react'
import { FireFilled } from '@ant-design/icons'

const Logo = () => {
  return (
    <div className='flex items-center justify-center text-white p-2.5'>
        <div className="w-[40px] h-[40px] flex items-center justify-center text-[1.5rem] rounded-[50%] bg-[#1890ff]">
            <FireFilled />
        </div>
    </div>
  )
}

export default Logo