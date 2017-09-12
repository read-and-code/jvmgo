package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_1
// Push the float constant <f> (1.0) onto the operand stack.
type FConst1 struct {
	base_instructions.NoOperandsInstruction
}

func (fConst1 *FConst1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(1.0)
}
