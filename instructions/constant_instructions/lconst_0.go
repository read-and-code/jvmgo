package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lconst_0
// Push the long constant <l> (0) onto the operand stack.
type LConst0 struct {
	base_instructions.NoOperandsInstruction
}

func (lConst0 *LConst0) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushLongValue(0)
}
