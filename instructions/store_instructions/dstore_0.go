package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_0
// Store double into local variable
type DStore0 struct {
	base_instructions.NoOperandsInstruction
}

func (dStore0 *DStore0) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 0)
}
