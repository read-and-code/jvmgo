package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup2
// Duplicate the top one or two values on the operand stack and push the duplicated value or values back onto the operand stack in the original order
type Dup2 struct {
	base_instructions.NoOperandsInstruction
}

func (dup2 *Dup2) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue1 := operandStack.PopOperand()
	operandValue2 := operandStack.PopOperand()

	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
	operandStack.PushOperand(operandValue2)
	operandStack.PushOperand(operandValue1)
}
