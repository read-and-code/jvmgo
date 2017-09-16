package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// i2l
// Convert int to long
type I2l struct {
	base_instructions.NoOperandsInstruction
}

func (i2l *I2l) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushLongValue(int64(integerValue))
}
