package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ldiv
// Divide long
type LDiv struct {
	base_instructions.NoOperandsInstruction
}

func (lDiv *LDiv) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue2 := operandStack.PopLongValue()
	longValue1 := operandStack.PopLongValue()

	operandStack.PushLongValue(longValue1 / longValue2)
}
