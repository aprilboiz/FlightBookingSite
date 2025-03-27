import React from "react";

const BusinessContext = () => {
  return (
    <div className="w-[700px] mx-auto text-lg">
      <h2 className="text-lg font-medium mt-5">BUSINESS CONTEXT</h2>
      <hr className="mb-5" />
      <p>
        Rùa Airline có kịch bản kinh doanh cơ bản như sau: <strong>Nhận lịch chuyến bay</strong>, <strong>Bán vé</strong>, <strong>Ghi nhận đặt vé</strong>, <strong>Tra cứu chuyến bay</strong>, <strong>Lập báo cáo tháng</strong>, <strong>Thay đổi quy định</strong>
      </p>
      <p><strong>Nhận lịch chuyến bay: </strong>người quản trị viên, nhân viên có thể xêm được lịc chuyến bay trong ngày hoặc tất cả các chuyến bay đang hoạt động. Trong chức năng này, người quản trị có thể thêm các chuyến bay, việc thêm các chuyến bay sẽ tuân thủ theo các qui định ngày bay không được thấp hơn ngày hiện tại, tuân thủ theo thời gian bay tối thiểu, số sân bay trung gian cũng như thời gian dừng khi tới sân bay trung gian</p>
      <p><strong>Bán vé: </strong>nhân viên có thể chọn chuyến bay, hoặc tìm kiếm chuyến bay theo mã chuyến bay, nhân viên sẽ điền đầy đủ thông tin khách hàng và nhấn đặt vé. Giá vé cũng sẽ được thay đổi tùy vào loại ghế được chọn</p>
      <p><strong>Ghi nhận đặt vé: </strong>hệ thống có thể quản lý phiếu đặt chỗ, nhân viên khi ghi nhận cần phải kiểm tra ngày đặt vé phải sơm hơn 1 ngày bay và nếu vào ngày khởi hành, các phiếu đặt vào ngày đó sẽ hủy</p>
      <p><strong>Tra cứu chuyến bay: </strong>hệ thống có thể quản lý việc tra cứu chuyến bay và các thông tin, nhân viên và khách hàng có thể tra cứu chuyến bay để biết chi tiết về các chuyến bay</p>
      <p><strong>Báo cáo doanh thu tháng: </strong>mỗi tháng, người quản trị sẽ lập ra phiếu báo cáo doanh thu tháng và gửi về cho giám đốc</p>
      <p><strong>Báo cáo doanh thu năm: </strong>mỗi năm, người quản trị sẽ lập ra phiếu báo cáo doanh thu tháng và gửi về cho giám đốc</p>
      <p><strong>Thay đổi quy định: </strong>người quản trị có thể thay đổi các quy định của hệ thống từ việc thay đổi số lượng sân bay nếu có sân bay mới, thời gian bay tối thiểu của mỗi chuyến bay, số sân bay trung gian tối đa khi di chuyển xa cũng như thời gian dừng tối thiểu ở các sân bay trung gian. Người quản trị còn có thể thay đổi về số lượng ghế của mỗi máy bay và cũng như số lượng hạng vé. Cuối cùng, người quản trị có thể thay đổi thời gian chậm nhất khi khách hàng đặt vé và thời gian hủy vé của khách hàng</p>
    </div>
  );
};

export default BusinessContext;
