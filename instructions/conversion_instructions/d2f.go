package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// d2f
// Convert double to float
type D2f struct {
	base_instructions.NoOperandsInstruction
}

func (d2f *D2f) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushFloatValue(float32(doubleValue))
}
