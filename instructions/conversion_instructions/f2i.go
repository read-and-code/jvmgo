package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// f2i
// Convert float to int
type F2i struct {
	base_instructions.NoOperandsInstruction
}

func (f2i *F2i) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushIntegerValue(int32(floatValue))
}
