package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dconst_0
type PushDoubleConstant0 struct {
	base_instructions.NoOperandsInstruction
}

func (pushDoubleConstant0 *PushDoubleConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushDoubleValue(0.0)
}
