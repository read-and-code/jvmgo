package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_1
type PushIntegerConstant1 struct {
	base.NoOperandsInstruction
}

func (pushIntegerConstant1 *PushIntegerConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(1)
}
