package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lconst_0
type PushLongConstant0 struct {
	base.NoOperandsInstruction
}

func (pushLongConstant0 *PushLongConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushLongValue(0)
}
