package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup2
type CopyPushOneOrTwoTopOperandStackValues struct {
	base_instructions.NoOperandsInstruction
}

func (copyPushOneOrTwoTopOperandStackValues *CopyPushOneOrTwoTopOperandStackValues) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue1 := operandStack.PopOperand()
	operandValue2 := operandStack.PopOperand()

	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
}
