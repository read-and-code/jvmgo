package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dup
// Duplicate the top value on the operand stack and push the duplicated value onto the operand stack
type Dup struct {
	base_instructions.NoOperandsInstruction
}

func (dup *Dup) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	operandValue := operandStack.PopOperand()

	operandStack.PushOperand(operandValue)
	operandStack.PushOperand(operandValue)
}
