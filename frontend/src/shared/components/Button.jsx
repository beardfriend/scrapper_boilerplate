import React from "react";
import { Button as ChakraButton } from "@chakra-ui/react";

function Button(props) {
  const { children } = props;
  return (
    <ChakraButton colorScheme="facebook" bg="black" color="#ffffff" borderRadius="20px" h="50px" {...props}>
      {children}
    </ChakraButton>
  );
}

export default Button;
