package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// return
// Return void from method
type Return struct {
	base_instructions.NoOperandsInstruction
}

func (returnInstruction *Return) Execute(frame *runtime_data_area.Frame) {
	frame.GetThread().PopFrame()
}
