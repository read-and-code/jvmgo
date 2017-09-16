package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2s
// Convert int to short
type I2s struct {
	base_instructions.NoOperandsInstruction
}

func (i2s *I2s) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(int32(int16(integerValue)))
}
