package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2c
// Convert int to char.
// The value on the top of the operand stack must be of type int.
// It is popped from the operand stack, truncated to char, then zero-extended to an int result.
// That result is pushed onto the operand stack.
type I2c struct {
	base_instructions.NoOperandsInstruction
}

func (i2c *I2c) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(int32(int16(integerValue)))
}
