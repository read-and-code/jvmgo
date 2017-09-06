package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_0
type PushIntegerConstant0 struct {
	base.NoOperandsInstruction
}

func (pushIntegerConstant0 *PushIntegerConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(0)
}
