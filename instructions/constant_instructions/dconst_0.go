package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dconst_0
// Push the double constant <d> (0.0) onto the operand stack
type DConst0 struct {
	base_instructions.NoOperandsInstruction
}

func (dConst0 *DConst0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushDoubleValue(0.0)
}
