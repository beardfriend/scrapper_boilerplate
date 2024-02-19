import React from "react";
import {
  NumberInput as ChakraNumberInput,
  NumberInputField,
  NumberInputStepper,
  NumberIncrementStepper,
  NumberDecrementStepper,
} from "@chakra-ui/react";

function NumberInput({ value, max, min, onChange, ...props }) {
  return (
    <ChakraNumberInput value={value} max={max} min={min} w="100px" onChange={onChange} {...props}>
      <NumberInputField />
      <NumberInputStepper>
        <NumberIncrementStepper />
        <NumberDecrementStepper />
      </NumberInputStepper>
    </ChakraNumberInput>
  );
}

export default NumberInput;
