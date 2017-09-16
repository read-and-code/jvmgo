package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fneg
// Negate float
type FNeg struct {
	base_instructions.NoOperandsInstruction
}

func (fNeg *FNeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushFloatValue(-floatValue)
}
