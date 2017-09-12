package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_1
// Push the int constant <i> (1) onto the operand stack.
type IConst1 struct {
	base_instructions.NoOperandsInstruction
}

func (iConst1 *IConst1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(1)
}
