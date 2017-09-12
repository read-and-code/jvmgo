package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore_2
// Store int into local variable.
// The <n> must be an index into the local variable array of the current frame.
// The value on the top of the operand stack must be of type int.
// It is popped from the operand stack, and the value of the local variable at <n> is set to value.
type IStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (iStore2 *IStore2) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, 2)
}
