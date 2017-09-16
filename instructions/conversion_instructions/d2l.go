package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// d2l
// Convert double to long
type D2l struct {
	base_instructions.NoOperandsInstruction
}

func (d2l *D2l) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushLongValue(int64(doubleValue))
}
