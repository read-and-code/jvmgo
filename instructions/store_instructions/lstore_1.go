package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_1
// Store long into local variable
type LStore1 struct {
	base_instructions.NoOperandsInstruction
}

func (lStore1 *LStore1) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 1)
}
