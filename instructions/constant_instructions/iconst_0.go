package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_0
// Push the int constant <i> (0) onto the operand stack.
type IConst0 struct {
	base_instructions.NoOperandsInstruction
}

func (iConst0 *IConst0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(0)
}
