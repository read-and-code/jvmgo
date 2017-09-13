package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lrem
// Remainder long.
// Both value1 and value2 must be of type long.
// The values are popped from the operand stack.
// The long result is value1 - (value1 / value2) * value2.
// The result is pushed onto the operand stack.
type LRem struct {
	base_instructions.NoOperandsInstruction
}

func (lRem *LRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue1 := operandStack.PopLongValue()
	longValue2 := operandStack.PopLongValue()

	if longValue2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushLongValue(longValue1 % longValue2)
}
