package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// l2d
// Convert long to double.
// The value on the top of the operand stack must be of type long.
// It is popped from the operand stack and converted to a double result using IEEE 754 round to nearest mode.
// The result is pushed onto the operand stack.
type L2d struct {
	base_instructions.NoOperandsInstruction
}

func (l2d *L2d) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()

	operandStack.PushDoubleValue(float64(longValue))
}
