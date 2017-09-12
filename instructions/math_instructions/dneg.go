package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dneg
// Negate double.
// The value must be of type double.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// The double result is the arithmetic negation of value'.
// The result is pushed onto the operand stack.
type DNeg struct {
	base_instructions.NoOperandsInstruction
}

func (dNeg *DNeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushDoubleValue(-doubleValue)
}
