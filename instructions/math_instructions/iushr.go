package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iushr
// Logical shift right int
type IUshr struct {
	base_instructions.NoOperandsInstruction
}

func (iUshr *IUshr) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue2 := operandStack.PopIntegerValue()
	integerValue1 := operandStack.PopIntegerValue()
	bitPositions := uint32(integerValue2) & 0x1f

	operandStack.PushIntegerValue(int32(uint32(integerValue1) >> bitPositions))
}
