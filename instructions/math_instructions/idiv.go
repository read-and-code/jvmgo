package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// idiv
// Divide int.
// Both value1 and value2 must be of type int.
// The values are popped from the operand stack.
// The int result is the value of the Java programming language expression value1 / value2.
// The result is pushed onto the operand stack.
type IDiv struct {
	base_instructions.NoOperandsInstruction
}

func (iDiv *IDiv) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue1 := operandStack.PopIntegerValue()
	integerValue2 := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(integerValue1 / integerValue2)
}
