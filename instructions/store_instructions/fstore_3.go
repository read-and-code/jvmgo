package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore_3
// Store float into local variable
type FStore3 struct {
	base_instructions.NoOperandsInstruction
}

func (fStore3 *FStore3) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, 3)
}
