package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore_0
// Store float into local variable.
// The <n> must be an index into the local variable array of the current frame.
// The value on the top of the operand stack must be of type float.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// The value of the local variable at <n> is set to value'.
type FStore0 struct {
	base_instructions.NoOperandsInstruction
}

func (fStore0 *FStore0) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, 0)
}
