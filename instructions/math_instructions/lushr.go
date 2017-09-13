package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lushr
// Logical shift right long.
// The value1 must be of type long, and value2 must be of type int.
// The values are popped from the operand stack.
// A long result is calculated by shifting value1 right logically (with zero extension) by the amount indicated by the low 6 bits of value2.
// The result is pushed onto the operand stack.
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
