package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dneg
// Negate double
type DNeg struct {
	base_instructions.NoOperandsInstruction
}

func (dNeg *DNeg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()

	operandStack.PushDoubleValue(-doubleValue)
}
