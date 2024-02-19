import { SelectDirectory } from "../../../wailsjs/go/directory/handler";
import { useState } from "react";

export const useDirectory = () => {
  const [isError, setIsError] = useState(false);
  const [path, setPath] = useState("");

  const handleDirectory = async () => {
    try {
      let p = await SelectDirectory();
      setPath(p);
    } catch (err) {
      setIsError(true);
    }
  };

  return { path, isError, handleDirectory };
};
