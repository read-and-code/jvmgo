package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lor
// Boolean OR long.
// Both value1 and value2 must be of type long.
// They are popped from the operand stack.
// A long result is calculated by taking the bitwise inclusive OR of value1 and value2.
// The result is pushed onto the operand stack.
type LOr struct {
	base_instructions.NoOperandsInstruction
}

func (lOr *LOr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue1 := operandStack.PopLongValue()
	longValue2 := operandStack.PopLongValue()

	operandStack.PushLongValue(longValue1 | longValue2)
}
