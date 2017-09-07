package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_1
type PushIntegerConstant1 struct {
	base_instructions.NoOperandsInstruction
}

func (pushIntegerConstant1 *PushIntegerConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(1)
}
