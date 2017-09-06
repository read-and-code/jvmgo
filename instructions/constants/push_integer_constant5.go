package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_5
type PushIntegerConstant5 struct {
	base.NoOperandsInstruction
}

func (pushIntegerConstant5 *PushIntegerConstant5) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(5)
}
