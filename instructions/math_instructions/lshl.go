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
	integerValue := operandStack.PopIntegerValue()
	longValue := operandStack.PopLongValue()
	bitPositions := uint32(integerValue) & 0x3f

	operandStack.PushLongValue(longValue << bitPositions)
}
