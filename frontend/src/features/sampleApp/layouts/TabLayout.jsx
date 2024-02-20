import React from "react";

import Header from "../../../shared/components/Header";
import { SimpleFooter } from "../../../shared/components/Footer";

function TabLayout({ children }) {
  return (
    <div className="w-100 mx-20">
      <Header />
      <div className="flex flex-col justify-center w-100">
        <div className="main mt-6">{children}</div>
      </div>
      <SimpleFooter />
    </div>
  );
}

export default TabLayout;
