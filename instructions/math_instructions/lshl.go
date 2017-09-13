package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lshl
// Shift left long.
// The value1 must be of type long, and value2 must be of type int.
// The values are popped from the operand stack.
// A long result is calculated by shifting value1 left by s bit positions,
// where s is the low 6 bits of value2.
// The result is pushed onto the operand stack.
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
