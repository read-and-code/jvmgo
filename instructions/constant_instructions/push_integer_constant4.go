package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iconst_4
type PushIntegerConstant4 struct {
	base_instructions.NoOperandsInstruction
}

func (pushIntegerConstant4 *PushIntegerConstant4) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushIntegerValue(4)
}
