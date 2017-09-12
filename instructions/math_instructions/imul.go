package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// imul
// Multiply int.
// Both value1 and value2 must be of type int.
// The values are popped from the operand stack.
// The int result is value1 * value2.
// The result is pushed onto the operand stack.
type IMul struct {
	base_instructions.NoOperandsInstruction
}

func (iMul *IMul) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue1 := operandStack.PopLongValue()
	longValue2 := operandStack.PopLongValue()

	operandStack.PushLongValue(longValue1 * longValue2)
}
