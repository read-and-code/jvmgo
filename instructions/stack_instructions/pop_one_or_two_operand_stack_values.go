package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// pop2
type PopOneOrTwoOperandStackValues struct {
	base_instructions.NoOperandsInstruction
}

func (popOneOrTwoOperandStackValues *PopOneOrTwoOperandStackValues) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()

	operandStack.PopOperand()
	operandStack.PopOperand()
}
