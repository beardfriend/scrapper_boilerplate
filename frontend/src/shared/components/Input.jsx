import React from "react";
import { Input as ChakraInput } from "@chakra-ui/react";

function Input(props) {
  return <ChakraInput border="1px solid black" variant="outline" h="50px" borderRadius="50px" {...props} />;
}

export default Input;
