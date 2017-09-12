package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ineg
// Negate int.
// The value must be of type int.
// It is popped from the operand stack.
// The int result is the arithmetic negation of value, -value.
// The result is pushed onto the operand stack.
type INeg struct {
	base_instructions.NoOperandsInstruction
}

func (iNeg *INeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(-integerValue)
}
