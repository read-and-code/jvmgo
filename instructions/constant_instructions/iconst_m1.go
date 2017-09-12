package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_m1
// Push the int constant <i> (-1) onto the operand stack.
type IConstM1 struct {
	base_instructions.NoOperandsInstruction
}

func (iConstM1 *IConstM1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(-1)
}
