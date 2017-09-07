package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_1
type PushFloatConstant1 struct {
	base_instructions.NoOperandsInstruction
}

func (pushFloatConstant1 *PushFloatConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(1.0)
}
