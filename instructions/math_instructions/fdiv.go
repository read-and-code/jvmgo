package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fdiv
// Divide float
type FDiv struct {
	base_instructions.NoOperandsInstruction
}

func (fDiv *FDiv) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue1 := operandStack.PopFloatValue()
	floatValue2 := operandStack.PopFloatValue()

	operandStack.PushFloatValue(floatValue1 / floatValue2)
}
