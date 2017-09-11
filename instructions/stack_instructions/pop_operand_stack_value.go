package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// pop
type PopOperandStackValue struct {
	base_instructions.NoOperandsInstruction
}

func (popOperandStackValue *PopOperandStackValue) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PopOperand()
}
