package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup2_x2
// Duplicate the top one or two values on the operand stack and insert the duplicated values, in the original order, into the operand stack
type Dup2X2 struct {
	base_instructions.NoOperandsInstruction
}

func (dup2X2 *Dup2X2) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue1 := operandStack.PopOperand()
	operandValue2 := operandStack.PopOperand()
	operandValue3 := operandStack.PopOperand()
	operandValue4 := operandStack.PopOperand()

	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
	operandStack.PushOperand(operandValue4)
	operandStack.PushOperand(operandValue3)
	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
}
