package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore_3
// Store int into local variable
type IStore3 struct {
	base_instructions.NoOperandsInstruction
}

func (iStore3 *IStore3) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, 3)
}
