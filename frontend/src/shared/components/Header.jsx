import React from "react";
import { Image } from "@nextui-org/react";
import Logo from "../../assets/logo/appicon.png";
import { useRecoilValue } from "recoil";
import { MetaState } from "../states/meta";

function Header() {
  const meta = useRecoilValue(MetaState);
  return (
    <header>
      <div className="flex items-center gap-10 p-10">
        <Image src={Logo} width={75} alt="loogo" />
        <h1 className="text-2xl font-bold">{meta?.name}</h1>
      </div>
    </header>
  );
}

export default Header;
