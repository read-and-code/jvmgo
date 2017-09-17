package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lshr
// Arithmetic shift right long
type LShr struct {
	base_instructions.NoOperandsInstruction
}

func (lShr *LShr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()
	longValue := operandStack.PopLongValue()
	bitPositions := uint32(integerValue) & 0x3f

	operandStack.PushLongValue(longValue >> bitPositions)
}
