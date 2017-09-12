package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_0
// Store long into local variable.
// Both <n> and <n>+1 must be indices into the local variable array of the current frame.
// The value on the top of the operand stack must be of type long.
// It is popped from the operand stack, and the local variables at <n> and <n>+1 are set to value.
type LStore0 struct {
	base_instructions.NoOperandsInstruction
}

func (lStore0 *LStore0) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 0)
}
