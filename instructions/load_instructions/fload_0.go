package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_0
// Load float from local variable.
// The <n> must be an index into the local variable array of the current frame.
// The local variable at <n> must contain a float.
// The value of the local variable at <n> is pushed onto the operand stack.
type FLoad0 struct {
	base_instructions.NoOperandsInstruction
}

func (fLoad0 *FLoad0) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 0)
}
