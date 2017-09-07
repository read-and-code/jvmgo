package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_3
type PushIntegerConstant3 struct {
	base_instructions.NoOperandsInstruction
}

func (pushIntegerConstant3 *PushIntegerConstant3) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(3)
}
