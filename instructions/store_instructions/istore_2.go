package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore_2
// Store int into local variable
type IStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (iStore2 *IStore2) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, 2)
}
