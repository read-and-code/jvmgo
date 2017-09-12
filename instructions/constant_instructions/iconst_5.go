package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_5
// Push the int constant <i> (5) onto the operand stack.
type IConst5 struct {
	base_instructions.NoOperandsInstruction
}

func (iConst5 *IConst5) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(5)
}
