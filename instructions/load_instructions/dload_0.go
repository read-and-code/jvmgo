package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload_0
// Load double from local variable
type DLoad0 struct {
	base_instructions.NoOperandsInstruction
}

func (dLoad0 *DLoad0) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, 0)
}
