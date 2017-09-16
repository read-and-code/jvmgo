package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// d2i
// Convert double to int
type D2i struct {
	base_instructions.NoOperandsInstruction
}

func (d2i *D2i) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushIntegerValue(int32(doubleValue))
}
