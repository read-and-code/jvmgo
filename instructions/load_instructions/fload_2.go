package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_2
// Load float from local variable
type FLoad2 struct {
	base_instructions.NoOperandsInstruction
}

func (fLoad2 *FLoad2) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 2)
}
