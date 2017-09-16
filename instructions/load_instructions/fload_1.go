package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_1
// Load float from local variable
type FLoad1 struct {
	base_instructions.NoOperandsInstruction
}

func (fLoad1 *FLoad1) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 1)
}
