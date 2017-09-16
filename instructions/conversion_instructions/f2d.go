package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// f2d
// Convert float to double
type F2d struct {
	base_instructions.NoOperandsInstruction
}

func (f2d *F2d) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushDoubleValue(float64(floatValue))
}
