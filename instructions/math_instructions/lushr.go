package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lushr
// Logical shift right long
type LUshr struct {
	base_instructions.NoOperandsInstruction
}

func (lUshr *LUshr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()
	longValue := operandStack.PopLongValue()
	bitPositions := uint32(integerValue) & 0x3f

	operandStack.PushLongValue(int64(uint64(longValue) >> bitPositions))
}
