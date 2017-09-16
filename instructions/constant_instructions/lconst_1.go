package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lconst_1
// Push the long constant <l> (1) onto the operand stack
type LConst1 struct {
	base_instructions.NoOperandsInstruction
}

func (lConst1 *LConst1) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushLongValue(1)
}
