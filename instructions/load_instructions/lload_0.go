package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_0
// Load long from local variable.
// Both <n> and <n>+1 must be indices into the local variable array of the current frame.
// The local variable at <n> must contain a long.
// The value of the local variable at <n> is pushed onto the operand stack.
type LLoad0 struct {
	base_instructions.NoOperandsInstruction
}

func (lLoad0 *LLoad0) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 0)
}
