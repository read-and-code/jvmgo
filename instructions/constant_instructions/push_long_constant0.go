package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lconst_0
type PushLongConstant0 struct {
	base_instructions.NoOperandsInstruction
}

func (pushLongConstant0 *PushLongConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushLongValue(0)
}
