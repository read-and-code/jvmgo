package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_0
// Load int from local variable
type ILoad0 struct {
	base_instructions.NoOperandsInstruction
}

func (iLoad0 *ILoad0) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 0)
}
