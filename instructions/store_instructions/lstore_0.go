package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_0
// Store long into local variable
type LStore0 struct {
	base_instructions.NoOperandsInstruction
}

func (lStore0 *LStore0) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 0)
}
