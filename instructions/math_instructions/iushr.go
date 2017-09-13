package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iushr
// Logical shift right int.
// Both value1 and value2 must be of type int.
// The values are popped from the operand stack.
// An int result is calculated by shifting value1 right by s bit positions,
// with zero extension, where s is the value of the low 5 bits of value2.
// The result is pushed onto the operand stack.
type IUshr struct {
	base_instructions.NoOperandsInstruction
}

func (iUshr *IUshr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue1 := operandStack.PopIntegerValue()
	integerValue2 := operandStack.PopIntegerValue()
	bitPositions := uint32(integerValue1) & 0x1f

	operandStack.PushIntegerValue(int32(uint32(integerValue2) >> bitPositions))
}
