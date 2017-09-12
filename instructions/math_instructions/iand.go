package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iand
// Boolean AND int.
// Both value1 and value2 must be of type int.
// They are popped from the operand stack.
// An int result is calculated by taking the bitwise AND (conjunction) of value1 and value2.
// The result is pushed onto the operand stack.
type IAnd struct {
	base_instructions.NoOperandsInstruction
}

func (iAnd *IAnd) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue1 := operandStack.PopIntegerValue()
	integerValue2 := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(integerValue1 & integerValue2)
}
