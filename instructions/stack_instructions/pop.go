package stack_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// pop
// Pop the top value from the operand stack
type Pop struct {
	base_instructions.NoOperandsInstruction
}

func (pop *Pop) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PopOperand()
}
