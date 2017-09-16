package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup_x2
// Duplicate the top value on the operand stack and insert the duplicated value two or three values down in the operand stack
type DupX2 struct {
	base_instructions.NoOperandsInstruction
}

func (dupX2 *DupX2) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue1 := operandStack.PopOperand()
	operandValue2 := operandStack.PopOperand()
	operandValue3 := operandStack.PopOperand()

	operandStack.PushOperand(operandValue1)
	operandStack.PushOperand(operandValue3)
	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
}
