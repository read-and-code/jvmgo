package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fconst_2
type PushFloatConstant2 struct {
	base_instructions.NoOperandsInstruction
}

func (pushFloatConstant2 *PushFloatConstant2) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushFloatValue(2.0)
}
