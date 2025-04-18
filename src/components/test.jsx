import { useState, useEffect } from "react";

const options = {
  root: [
    { label: "Cách thêm chuyến bay", value: "addplane" },
    { label: "Chỉnh sửa / xóa chuyến bay", value: "updateordeleteplane" },
    // { label: "Cập nhật lịch chuyến", value: "lichchuyen" },
  ],
};

const responses = {
  addplane: 'Bạn chọn vào mục "Lịch chuyến bay" → Nhập thông tin chuyến bay → bấm vào nút "Thêm chuyến bay".',
  updateordeleteplane: 'Bạn chọn vào mục "Danh sách chuyến bay" → Bấm vào biểu tượng edit/xóa ở cuối mỗi dòng → Chỉnh sửa thông tin hoặc bấm biểu tượng nút "Xóa".',
//   lichchuyen:
//     'Vào module "Lịch chuyến" → Chọn chuyến bay → Chỉnh sửa giờ hoặc ngày → Lưu.',
};

function TypingBubble() {
  return (
    <div className="bg-blue-100 text-blue-900 self-start rounded-2xl px-3 py-2 text-sm max-w-[80%] animate-pulse">
      💬 Đang soạn...
    </div>
  );
}

export default function Test() {
  const [messages, setMessages] = useState([
    { from: "bot", text: "Xin chào! Tôi là chatbot được cấu hình sẵn các câu hỏi, tôi ở đây để giúp các bạn lần đầu tiếp xúc hệ thống" },
    { from: "bot", text: "Các bạn muối hỏi tôi điều gì?" },
  ]);
  const [isTyping, setIsTyping] = useState(false);
  const [open, setOpen] = useState(false);

  const handleClick = (value, label) => {
    const reply = responses[value];

    setMessages((prev) => [...prev, { from: "user", text: label }]);

    setIsTyping(true);

    setTimeout(() => {
      setMessages((prev) => [...prev, { from: "bot", text: "" }]);
    }, 800);

    setTimeout(() => {
      setMessages((prev) => {
        const newMsgs = [...prev];
        newMsgs[newMsgs.length - 1] = { from: "bot", text: reply };
        return newMsgs;
      });
      setIsTyping(false);
    }, 2000);
  };

  return (
    <div className="fixed bottom-4 right-4 z-50">
      {!open && (
        <button
          onClick={() => setOpen(true)}
          className="bg-blue-600 hover:bg-blue-700 text-white rounded-full w-14 h-14 shadow-lg flex items-center justify-center text-2xl"
        >
          💬
        </button>
      )}

      {open && (
        <div className="bg-white shadow-xl rounded-2xl p-4 w-80 h-[600px] flex flex-col">
          <div className="flex items-center justify-between mb-2">
            <div className="flex items-center space-x-2">
              <div className="w-12 h-12 rounded-full overflow-hidden mt-1">
                <img
                  src="https://img.freepik.com/premium-photo/capybara-logo_508233-163.jpg"
                  alt=""
                />
              </div>
              <h1 className="text-lg font-bold">Trợ lý Capybara</h1>
            </div>
            <button
              onClick={() => setOpen(false)}
              className="text-gray-500 hover:text-black"
            >
              ✖
            </button>
          </div>

          <div className="flex-1 overflow-y-auto space-y-2 p-2 border rounded-xl bg-gray-50">
            {messages.map((msg, i) => (
              <div
                key={i}
                className={`max-w-[80%] p-3 rounded-2xl text-sm whitespace-pre-line ${
                  msg.from === "bot"
                    ? "bg-blue-100 text-blue-900 self-start"
                    : "bg-green-200 text-green-900 self-end"
                }`}
              >
                {msg.text}
              </div>
            ))}
            {isTyping && <TypingBubble />}
          </div>

          <div className="mt-4 grid grid-cols-1 gap-2">
            {options.root.map((opt) => (
              <button
                key={opt.value}
                onClick={() => handleClick(opt.value, opt.label)}
                disabled={isTyping}
                className={`rounded-xl py-2 px-4 text-sm transition-all duration-200 ${
                  isTyping
                    ? "bg-gray-300 text-gray-500 cursor-not-allowed"
                    : "bg-blue-500 hover:bg-blue-600 text-white"
                }`}
              >
                {opt.label}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
