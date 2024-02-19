import { Card } from "@nextui-org/react";
import React from "react";

function Main({ children }) {
  return <Card className="tab_card">{children}</Card>;
}

export default Main;
