package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fneg
// Negate float.
// The value must be of type float.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// The float result is the arithmetic negation of value'.
// This result is pushed onto the operand stack.
type FNeg struct {
	base_instructions.NoOperandsInstruction
}

func (fNeg *FNeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()

	operandStack.PushFloatValue(-floatValue)
}
