import { Card } from "@nextui-org/react";
import React from "react";

function Main({ children }) {
  return <Card className="min-h-160 w-100 mt-3 p-6">{children}</Card>;
}

export default Main;
