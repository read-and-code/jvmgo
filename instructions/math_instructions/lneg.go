package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lneg
// Negate long
type LNeg struct {
	base_instructions.NoOperandsInstruction
}

func (lNeg *LNeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()

	operandStack.PushLongValue(-longValue)
}
