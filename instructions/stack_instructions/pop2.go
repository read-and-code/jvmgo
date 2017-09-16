package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// pop2
// Pop the top one or two values from the operand stack
type Pop2 struct {
	base_instructions.NoOperandsInstruction
}

func (pop2 *Pop2) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()

	operandStack.PopOperand()
	operandStack.PopOperand()
}
