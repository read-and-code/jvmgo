package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// l2f
// Convert long to float.
// The value on the top of the operand stack must be of type long.
// It is popped from the operand stack and converted to a float result using IEEE 754 round to nearest mode.
// The result is pushed onto the operand stack.
type L2f struct {
	base_instructions.NoOperandsInstruction
}

func (l2f *L2f) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()

	operandStack.PushFloatValue(float32(longValue))
}
