import React from "react";

const UserStory = () => {
  return (
    <div className="w-[700px] mx-auto text-lg">
      <h2 className="text-lg font-medium mt-5">USER STORY</h2>
      <hr className="mb-5" />
      <h3 className="text-lg font-medium mt-5">Nhận lịch chuyến bay</h3>
      <ul className="list-disc ml-5">
        <li>
          Là người quản trị, nhân viên tôi muốn xem được lịch chuyến bay trong ngày hoặc tất cả các chuyến bay đang hoạt động để biết được tình hình chuyến bay
        </li>
        <li>
          Là người quản trị tôi muốn thêm các chuyến bay mới để cập nhật lịch bay
        </li>
      </ul>
      <h3 className="text-lg font-medium mt-5">Bán vé</h3>
      <ul className="list-disc ml-5">
        <li>
          Là nhân viên, tôi muốn chọn chuyến bay, tìm kiếm chuyến bay theo mã chuyến bay để bán vé
        </li>
      </ul>
      <h3 className="text-lg font-medium mt-5">Ghi nhận đặt vé</h3>
      <ul className="list-disc ml-5">
        <li>
          Là nhân viên, tôi muốn ghi nhận đặt vé, kiểm tra ngày đặt vé phải sơm hơn 1 ngày bay và nếu vào ngày khởi hành, các phiếu đặt vào ngày đó sẽ hủy
        </li>
      </ul>
      <h3 className="text-lg font-medium mt-5">Tra cứu chuyến bay</h3>
      <ul className="list-disc ml-5">
        <li>
          Là nhân viên, tôi muốn tra cứu chuyến bay để biết các thông tin, và kiểm tra tình hình chuyến bay
        </li>
        <li>
          Là khách hàng, tôi muốn tra cứu chuyến bay để biết thông tin về chuyến bay
        </li>
      </ul>
      <h3 className="text-lg font-medium mt-5">Lập báo cáo tháng</h3>
      <ul className="list-disc ml-5">
        <li>
          Là người quản trị, tôi muốn lập báo cáo doang thu hằng tháng và gửi về cho giám đốc
        </li>
      </ul>
      <h3 className="text-lg font-medium mt-5">Lập báo cáo năm</h3>
      <ul className="list-disc ml-5">
        <li>
          Là người quản trị, tôi muốn lập báo cáo doang thu hằng năm và gửi về cho giám đốc
        </li>
      </ul>
      <h3 className="text-lg font-medium mt-5">Thay đổi quy định</h3>
      <ul className="list-disc ml-5">
        <li>
          Là người quản trị, tôi muốn thay đổi các quy định của hệ thống từ việc thay đổi số lượng sân bay nếu có sân bay mới, thời gian bay tối thiểu của mỗi chuyến bay, số sân bay trung gian tối đa khi di chuyển xa cũng như thời gian dừng tối thiểu ở các sân bay trung gian
        </li>
      </ul>
    </div>
  );
};

export default UserStory;
