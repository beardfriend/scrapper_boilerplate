import { useEffect, useRef } from "react";
export function useOnlyFisrtRender(func) {
  const firstRender = useRef(false);
  useEffect(() => {
    if (firstRender.current) {
      return;
    }
    func();

    firstRender.current = true;
  }, []);
}
