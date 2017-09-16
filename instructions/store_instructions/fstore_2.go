package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore_2
// Store float into local variable
type FStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (fStore2 *FStore2) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, 2)
}
