package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_3
// Load int from local variable.
// The <n> must be an index into the local variable array of the current frame.
// The local variable at <n> must contain an int.
// The value of the local variable at <n> is pushed onto the operand stack.
type ILoad3 struct {
	base_instructions.NoOperandsInstruction
}

func (iLoad3 *ILoad3) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 3)
}
