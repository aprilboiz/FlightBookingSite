import React from "react";
import ConceptualModelImage from "../assets/ConceptualModel.png";

const ConceptualModel = () => {
  const downloadImage = () => {
    const link = document.createElement("a");
    link.href = ConceptualModelImage;
    link.download = "ConceptualModel.png";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };
  return (
    <div className="w-[700px] mx-auto text-lg">
      <h2 className="text-lg font-medium mt-5">BUSINESS CONTEXT</h2>
      <hr className="mb-5" />
      <img
        className="w-full mb-4"
        src={ConceptualModelImage}
        alt="Conceptual Model"
      />
      <div className="text-sm text-right">
        <a
          href="/ConceptualModelDraw.drawio"
          download="ConceptualModelDraw.drawio"
          className="px-4 py-2 text-blue-500"
        >
          Tải xuống file Draw.io
        </a>
        <button
          onClick={downloadImage}
          className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 transition"
        >
          Download Image
        </button>
      </div>
    </div>
  );
};

export default ConceptualModel;
