package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_3
// Push the int constant <i> (3) onto the operand stack.
type IConst3 struct {
	base_instructions.NoOperandsInstruction
}

func (iConst3 *IConst3) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(3)
}
