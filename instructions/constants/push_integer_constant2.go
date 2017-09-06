package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_2
type PushIntegerConstant2 struct {
	base.NoOperandsInstruction
}

func (pushIntegerConstant2 *PushIntegerConstant2) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(2)
}
