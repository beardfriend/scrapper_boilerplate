import React from "react";
import { useSetRecoilState } from "recoil";
import { MetaState } from "../states/meta";
import { GetAppMeta } from "../../../wailsjs/go/meta/handler";
import { useOnlyFisrtRender } from "../hooks/useOnlyFirstRender";

export default function MetaLayout({ children }) {
  const setMeta = useSetRecoilState(MetaState);
  useOnlyFisrtRender(() => {
    GetAppMeta().then((res) => {
      setMeta(res);
    });
  });

  return <>{children}</>;

  // 인증이 꼭 필요한 경우
}
