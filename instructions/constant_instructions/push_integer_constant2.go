package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_2
type PushIntegerConstant2 struct {
	base_instructions.NoOperandsInstruction
}

func (pushIntegerConstant2 *PushIntegerConstant2) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(2)
}
