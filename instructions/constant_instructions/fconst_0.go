package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_0
// Push the float constant <f> (0.0) onto the operand stack.
type FConst0 struct {
	base_instructions.NoOperandsInstruction
}

func (fConst0 *FConst0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(0.0)
}
