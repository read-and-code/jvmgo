package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dconst_1
type PushDoubleConstant1 struct {
	base_instructions.NoOperandsInstruction
}

func (pushDoubleConstant1 *PushDoubleConstant1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushDoubleValue(1.0)
}
