package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// f2l
// Convert float to long
type F2l struct {
	base_instructions.NoOperandsInstruction
}

func (f2l *F2l) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushLongValue(int64(floatValue))
}
