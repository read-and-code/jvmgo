package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dconst_0
type PushDoubleConstant0 struct {
	base.NoOperandsInstruction
}

func (pushDoubleConstant0 *PushDoubleConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushDoubleValue(0.0)
}
