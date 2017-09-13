package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2f
// Convert int to float.
// The value on the top of the operand stack must be of type int.
// It is popped from the operand stack and converted to the float result using IEEE 754 round to nearest mode.
// The result is pushed onto the operand stack.
type I2f struct {
	base_instructions.NoOperandsInstruction
}

func (i2f *I2f) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushFloatValue(float32(integerValue))
}
