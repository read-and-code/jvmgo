package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// irem
// Remainder int.
// Both value1 and value2 must be of type int.
// The values are popped from the operand stack.
// The int result is value1 - (value1 / value2) * value2.
// The result is pushed onto the operand stack.
type IRem struct {
	base_instructions.NoOperandsInstruction
}

func (iRem *IRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue1 := operandStack.PopIntegerValue()
	integerValue2 := operandStack.PopIntegerValue()

	if integerValue2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushIntegerValue(integerValue1 % integerValue2)
}
