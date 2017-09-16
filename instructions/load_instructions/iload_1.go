package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_1
// Load int from local variable
type ILoad1 struct {
	base_instructions.NoOperandsInstruction
}

func (iLoad1 *ILoad1) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 1)
}
