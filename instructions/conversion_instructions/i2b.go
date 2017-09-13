package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2b
// Convert int to byte.
// The value on the top of the operand stack must be of type int.
// It is popped from the operand stack, truncated to a byte, then sign-extended to an int result.
// That result is pushed onto the operand stack.
type I2b struct {
	base_instructions.NoOperandsInstruction
}

func (i2b *I2b) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(int32(int8(integerValue)))
}
