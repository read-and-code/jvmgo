package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_m1
type PushIntegerConstantNegative1 struct {
	base.NoOperandsInstruction
}

func (pushIntegerConstantNegative1 *PushIntegerConstantNegative1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(-1)
}
