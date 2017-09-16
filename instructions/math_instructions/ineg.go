package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ineg
// Negate int
type INeg struct {
	base_instructions.NoOperandsInstruction
}

func (iNeg *INeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(-integerValue)
}
