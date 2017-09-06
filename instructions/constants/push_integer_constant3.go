package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_3
type PushIntegerConstant3 struct {
	base.NoOperandsInstruction
}

func (pushIntegerConstant3 *PushIntegerConstant3) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(3)
}
