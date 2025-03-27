import React from 'react'

const Website = () => {
  return (
    <div className='w-[700px] mx-auto text-lg'>
        <h2 className='text-lg font-medium mt-5'>LÝ DO CHỌN ĐỀ TÀI</h2>
        <hr className='mb-5'/>
        <p>Dưới sự phát triển không ngừng của khoa học và công nghệ, đặc biệt trong lĩnh vực giao thông vận tải, các phương tiện di chuyển trong tương lai sẽ ngày càng hiện đại hơn, tiết kiệm nhiên liệu và thân thiện với môi trường.</p>
        <p>Hiện nay, Việt Nam có nhiều hãng hàng không nội địa phục vụ hành khách, nổi bật là Vietnam Airlines, Vietjet Air, Jetstar Pacific và Bamboo Airways. Trước đây, khi chỉ có Vietnam Airlines hoạt động với mức giá cao, nhiều người thường lựa chọn tàu hỏa hoặc xe khách cho các chuyến đi xa. Tuy nhiên, sự cạnh tranh giữa các hãng hàng không đã giúp giá vé máy bay trở nên hợp lý hơn, tạo điều kiện cho du khách di chuyển thuận tiện với chi phí thấp.</p>
        <p>Cùng với sự mở rộng và phát triển của ngành hàng không, nhu cầu đặt vé máy bay ngày càng tăng cao. Để tối ưu hóa quy trình bán vé, giảm thiểu chi phí và nâng cao trải nghiệm cho khách hàng cũng như doanh nghiệp, việc xây dựng một phần mềm quản lý bán vé chuyến bay là rất cần thiết.</p>
        <p>Xuất phát từ thực tế này, nhóm chúng em quyết định lựa chọn đề tài nghiên cứu về hệ thống quản lý bán vé máy bay. Với những kiến thức và kỹ năng sẵn có, nhóm mong muốn tìm hiểu và phát triển một giải pháp tối ưu, góp phần nâng cao hiệu quả quản lý và mang lại sự tiện lợi cho người sử dụng.</p>

        <h2 className='text-lg font-medium mt-5'>RÙA AIRLINE WEBSITE</h2>
        <hr className='mb-5'/>
        <p>Rùa Airline Website là một nền tảng trực tuyến chuyên cung cấp dịch vụ đặt vé máy bay cho khách hàng, giúp họ dễ dàng tìm kiếm, so sánh và đặt chỗ các chuyến bay một cách nhanh chóng, thuận tiện. Trang web được xây dựng trên nền tảng công nghệ hiện đại, với Spring Boot làm backend và Vite + React để phát triển giao diện người dùng, đảm bảo hiệu suất cao và trải nghiệm mượt mà.</p>

        <h2 className='text-lg font-medium mt-5'>CÔNG NGHỆ SỬ DỤNG</h2>
        <hr className='mb-5'/>
        <ul className='list-disc pl-5'>
            <li><strong>Backend:</strong> Hệ thống sử dụng <strong><a href="https://spring.io/projects/spring-boot" className='text-blue-500'>Spring Boot</a></strong>, một framework mạnh mẽ của Java, giúp xây dựng các API RESTful một cách linh hoạt, bảo mật cao và có khả năng mở rộng tốt. Backend đóng vai trò quan trọng trong việc xử lý dữ liệu, xác thực người dùng, cũng như đảm bảo tốc độ phản hồi nhanh chóng khi có yêu cầu từ phía khách hàng.</li>
            <li><strong>Cơ sở dữ liệu:</strong> Dữ liệu được lưu trữ trên <strong><a href="" className='text-blue-500'>PostgreSQL</a></strong>, một hệ quản trị cơ sở dữ liệu quan hệ mã nguồn mở phổ biến, cung cấp hiệu suất cao, khả năng mở rộng và độ tin cậy tốt. PostgreSQL hỗ trợ xử lý lượng dữ liệu lớn, đảm bảo tính toàn vẹn dữ liệu cho hệ thống đặt vé.</li>
            <li><strong>Frontend:</strong> Giao diện người dùng được phát triển bằng React.js kết hợp với Vite, giúp tăng tốc quá trình xây dựng ứng dụng và tối ưu hiệu suất hiển thị. React.js mang đến khả năng tái sử dụng component, giúp giao diện web trực quan, dễ dàng mở rộng và bảo trì.</li>
        </ul>
    </div>
  )
}

export default Website