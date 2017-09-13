package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// d2i
// Convert double to int.
// The value on the top of the operand stack must be of type double.
// It is popped from the operand stack and undergoes value set conversion resulting in value'.
// Then value' is converted to an int.
// The result is pushed onto the operand stack.
type D2i struct {
	base_instructions.NoOperandsInstruction
}

func (d2i *D2i) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushIntegerValue(int32(doubleValue))
}
