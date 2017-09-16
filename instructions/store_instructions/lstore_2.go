package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_2
// Store long into local variable
type LStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (lStore2 *LStore2) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 2)
}
