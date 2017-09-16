package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2d
// Convert int to double
type I2d struct {
	base_instructions.NoOperandsInstruction
}

func (i2d *I2d) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushDoubleValue(float64(integerValue))
}
