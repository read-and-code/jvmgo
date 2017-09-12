package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_4
// Push the int constant <i> (4) onto the operand stack.
type IConst4 struct {
	base_instructions.NoOperandsInstruction
}

func (iConst4 *IConst4) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(4)
}
