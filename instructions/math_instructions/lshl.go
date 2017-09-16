package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lshl
// Shift left long
type LShl struct {
	base_instructions.NoOperandsInstruction
}

func (lShl *LShl) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue1 := operandStack.PopLongValue()
	longValue2 := operandStack.PopLongValue()
	bitPositions := uint32(longValue2) & 0x3f

	operandStack.PushLongValue(longValue1 << bitPositions)
}
