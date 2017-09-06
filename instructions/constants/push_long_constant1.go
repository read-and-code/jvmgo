package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lconst_1
type PushLongConstant1 struct {
	base.NoOperandsInstruction
}

func (pushLongConstant1 *PushLongConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushLongValue(1)
}
