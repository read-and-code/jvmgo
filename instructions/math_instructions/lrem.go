package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lrem
// Remainder long
type LRem struct {
	base_instructions.NoOperandsInstruction
}

func (lRem *LRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue2 := operandStack.PopLongValue()
	longValue1 := operandStack.PopLongValue()

	if longValue2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushLongValue(longValue1 % longValue2)
}
