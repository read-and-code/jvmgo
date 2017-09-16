package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// nop
// Do nothing
type Nop struct {
	base_instructions.NoOperandsInstruction
}

func (nop *Nop) Execute(frame *runtime_data_area.Frame) {
}
