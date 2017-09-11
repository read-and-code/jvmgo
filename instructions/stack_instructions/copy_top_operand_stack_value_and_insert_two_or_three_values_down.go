package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup_x2
type CopyTopOperandStackValueAndInsertTwoOrThreeValuesDown struct {
	base_instructions.NoOperandsInstruction
}

func (copyTopOperandStackValueAndInsertTwoOrThreeValuesDown *CopyTopOperandStackValueAndInsertTwoOrThreeValuesDown) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue1 := operandStack.PopOperand()
	operandValue2 := operandStack.PopOperand()
	operandValue3 := operandStack.PopOperand()

	operandStack.PushOperand(operandValue1)
	operandStack.PushOperand(operandValue3)
	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
}
