package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dconst_1
// Push the double constant <d> (1.0) onto the operand stack
type DConst1 struct {
	base_instructions.NoOperandsInstruction
}

func (dConst1 *DConst1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushDoubleValue(1.0)
}
