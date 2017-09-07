package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_0
type PushFloatConstant0 struct {
	base_instructions.NoOperandsInstruction
}

func (pushFloatConstant0 *PushFloatConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(0.0)
}
