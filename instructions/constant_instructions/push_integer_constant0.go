package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_0
type PushIntegerConstant0 struct {
	base_instructions.NoOperandsInstruction
}

func (pushIntegerConstant0 *PushIntegerConstant0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(0)
}
