import React from "react";
import { Select as ChakraSelect } from "@chakra-ui/react";

function Select(props) {
  const { children } = props;
  return (
    <ChakraSelect
      borderRadius="50px"
      variant="outline"
      border="1px solid black"
      h="50px"
      colorScheme="facebook"
      {...props}
    >
      {children}
    </ChakraSelect>
  );
}

export default Select;
