import React from "react";

function TabLayout({ children }) {
  return (
    <div className="w-100 mx-20">
      <div className="flex flex-col justify-center w-100">
        <div className="main mt-6">{children}</div>
      </div>
      <footer className="flex flex-col justify-center items-center bottom-0 h-32 min-w-100 bg-gray-700 text-white">
        <div className="text-gray-400">@2024 All right reserved MINDFLEX</div>
        <p className="text-center mt-3">
          본 프로그램을 스스로 복제, 개작, 번역, 배포, 발송 및 전송할 수 없으며
          <br /> 타인에게 이러한 권리를 허락할 수 있는 대여권 또한 금지합니다.
          <br /> 이를 어길 시 법적 처벌을 받을 수 있습니다.
        </p>
      </footer>
    </div>
  );
}

export default TabLayout;
