package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_1
type PushFloatConstant1 struct {
	base.NoOperandsInstruction
}

func (pushFloatConstant1 *PushFloatConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(1.0)
}
