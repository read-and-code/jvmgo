package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_3
// Load float from local variable
type FLoad3 struct {
	base_instructions.NoOperandsInstruction
}

func (fLoad3 *FLoad3) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 3)
}
