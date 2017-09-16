package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ior
// Boolean OR int
type IOr struct {
	base_instructions.NoOperandsInstruction
}

func (iOr *IOr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue1 := operandStack.PopIntegerValue()
	integerValue2 := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(integerValue1 | integerValue2)
}
