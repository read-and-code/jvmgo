package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2f
// Convert int to float
type I2f struct {
	base_instructions.NoOperandsInstruction
}

func (i2f *I2f) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushFloatValue(float32(integerValue))
}
