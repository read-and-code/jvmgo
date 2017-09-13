package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// d2f
// Convert double to float.
// The value on the top of the operand stack must be of type double.
// It is popped from the operand stack and undergoes value set conversion resulting in value'.
// Then value' is converted to a float result using IEEE 754 round to nearest mode.
// The result is pushed onto the operand stack.
type D2f struct {
	base_instructions.NoOperandsInstruction
}

func (d2f *D2f) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushFloatValue(float32(doubleValue))
}
