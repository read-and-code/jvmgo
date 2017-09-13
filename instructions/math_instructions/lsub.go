package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lsub
// Subtract long.
// Both value1 and value2 must be of type long.
// The values are popped from the operand stack.
// The long result is value1 - value2.
// The result is pushed onto the operand stack.
type LSub struct {
	base_instructions.NoOperandsInstruction
}

func (lSub *LSub) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue1 := operandStack.PopLongValue()
	longValue2 := operandStack.PopLongValue()

	operandStack.PushLongValue(longValue1 - longValue2)
}
