package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// irem
// Remainder int
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
