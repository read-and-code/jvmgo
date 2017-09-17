package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ishl
// Shift left int
type IShl struct {
	base_instructions.NoOperandsInstruction
}

func (iShl *IShl) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue2 := operandStack.PopIntegerValue()
	integerValue1 := operandStack.PopIntegerValue()
	bitPositions := uint32(integerValue2) & 0x1f

	operandStack.PushIntegerValue(integerValue1 << bitPositions)
}
