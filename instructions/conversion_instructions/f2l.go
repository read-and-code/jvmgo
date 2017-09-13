package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// f2l
// Convert float to long.
// The value on the top of the operand stack must be of type float.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// Then value' is converted to a long result.
// This result is pushed onto the operand stack.
type F2l struct {
	base_instructions.NoOperandsInstruction
}

func (f2l *F2l) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushLongValue(int64(floatValue))
}
