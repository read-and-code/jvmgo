package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// nop
type NoOperation struct {
	base_instructions.NoOperandsInstruction
}

func (noOperation *NoOperation) Execute(frame *runtime_data_area.Frame) {
}
