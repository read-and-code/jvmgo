package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2d
// Convert int to double.
// The value on the top of the operand stack must be of type int.
// It is popped from the operand stack and converted to a double result.
// The result is pushed onto the operand stack.
type I2d struct {
	base_instructions.NoOperandsInstruction
}

func (i2d *I2d) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushDoubleValue(float64(integerValue))
}
