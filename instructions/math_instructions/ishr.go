package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ishr
// Arithmetic shift right int.
// Both value1 and value2 must be of type int.
// The values are popped from the operand stack.
// An int result is calculated by shifting value1 right by s bit positions,
// with sign extension, where s is the value of the low 5 bits of value2.
// The result is pushed onto the operand stack.
type IShr struct {
	base_instructions.NoOperandsInstruction
}

func (iShr *IShr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue1 := operandStack.PopIntegerValue()
	integerValue2 := operandStack.PopIntegerValue()
	bitPositions := uint32(integerValue2) & 0x1f

	operandStack.PushIntegerValue(integerValue1 >> bitPositions)
}
