package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_0
type PushFloatConstant0 struct {
	base.NoOperandsInstruction
}

func (pushFloatConstant0 *PushFloatConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(0.0)
}
