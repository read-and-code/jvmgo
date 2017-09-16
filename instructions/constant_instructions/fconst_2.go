package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_2
// Push the float constant <f> (2.0) onto the operand stack
type FConst2 struct {
	base_instructions.NoOperandsInstruction
}

func (fConst2 *FConst2) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(2.0)
}
