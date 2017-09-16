package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dadd
// Add double
type DAdd struct {
	base_instructions.NoOperandsInstruction
}

func (dAdd *DAdd) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue1 := operandStack.PopDoubleValue()
	doubleValue2 := operandStack.PopDoubleValue()

	operandStack.PushDoubleValue(doubleValue1 + doubleValue2)
}
