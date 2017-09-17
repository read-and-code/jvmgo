package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// isub
// Subtract int
type ISub struct {
	base_instructions.NoOperandsInstruction
}

func (iSub *ISub) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue2 := operandStack.PopIntegerValue()
	integerValue1 := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(integerValue1 - integerValue2)
}
