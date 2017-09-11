package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup
type CopyPushTopOperandStackValue struct {
	base_instructions.NoOperandsInstruction
}

func (copyPushTopOperandStackValue *CopyPushTopOperandStackValue) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue := operandStack.PopOperand()

	operandStack.PushOperand(operandValue)
	operandStack.PushOperand(operandValue)
}
