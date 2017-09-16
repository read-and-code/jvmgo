package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_2
// Store double into local variable
type DStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (dStore2 *DStore2) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 2)
}
