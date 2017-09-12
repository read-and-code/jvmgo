package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lneg
// Negate long.
// The value must be of type long.
// It is popped from the operand stack.
// The long result is the arithmetic negation of value, -value.
// The result is pushed onto the operand stack.
type LNeg struct {
	base_instructions.NoOperandsInstruction
}

func (lNeg *LNeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()

	operandStack.PushLongValue(-longValue)
}
