package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_3
// Store double into local variable
type DStore3 struct {
	base_instructions.NoOperandsInstruction
}

func (dStore3 *DStore3) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 3)
}
