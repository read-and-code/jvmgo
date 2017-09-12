package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_3
// Store double into local variable.
// Both <n> and <n>+1 must be indices into the local variable array of the current frame.
// The value on the top of the operand stack must be of type double.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// The local variables at <n> and <n>+1 are set to value'.
type DStore3 struct {
	base_instructions.NoOperandsInstruction
}

func (dStore3 *DStore3) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 3)
}
