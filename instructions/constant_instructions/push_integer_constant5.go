package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_5
type PushIntegerConstant5 struct {
	base_instructions.NoOperandsInstruction
}

func (pushIntegerConstant5 *PushIntegerConstant5) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(5)
}
