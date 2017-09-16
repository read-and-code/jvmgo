package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fsub
// Subtract float
type FSub struct {
	base_instructions.NoOperandsInstruction
}

func (fSub *FSub) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue1 := operandStack.PopFloatValue()
	floatValue2 := operandStack.PopFloatValue()

	operandStack.PushFloatValue(floatValue1 - floatValue2)
}
