package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lconst_1
type PushLongConstant1 struct {
	base_instructions.NoOperandsInstruction
}

func (pushLongConstant1 *PushLongConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushLongValue(1)
}
