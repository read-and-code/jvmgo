package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload
// Load long from local variable.
// The index is an unsigned byte.
// Both index and index+1 must be indices into the local variable array of the current frame.
// The local variable at index must contain a long.
// The value of the local variable at index is pushed onto the operand stack.
type LLoad struct {
	base_instructions.Index8Instruction
}

func (lLoad *LLoad) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, uint(lLoad.Index))
}

func loadLongValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetLongValue(index)

	frame.GetOperandStack().PushLongValue(value)
}
