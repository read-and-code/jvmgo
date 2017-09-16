package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload_1
// Load double from local variable
type DLoad1 struct {
	base_instructions.NoOperandsInstruction
}

func (dLoad1 *DLoad1) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, 1)
}
