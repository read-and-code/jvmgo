package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// f2d
// Convert float to double.
// The value on the top of the operand stack must be of type float.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// Then value' is converted to a double result.
// This result is pushed onto the operand stack.
type F2d struct {
	base_instructions.NoOperandsInstruction
}

func (f2d *F2d) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushDoubleValue(float64(floatValue))
}
