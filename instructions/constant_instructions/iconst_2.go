package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_2
// Push the int constant <i> (2) onto the operand stack.
type IConst2 struct {
	base_instructions.NoOperandsInstruction
}

func (iConst2 *IConst2) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(2)
}
