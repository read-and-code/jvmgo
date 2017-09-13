package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// l2i
// Convert long to int.
// The value on the top of the operand stack must be of type long.
// It is popped from the operand stack and converted to an int result
// by taking the low-order 32 bits of the long value and discarding the high-order 32 bits.
// The result is pushed onto the operand stack.
type L2i struct {
	base_instructions.NoOperandsInstruction
}

func (l2i *L2i) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()

	operandStack.PushIntegerValue(int32(longValue))
}
