package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_3
// Load int from local variable
type ILoad3 struct {
	base_instructions.NoOperandsInstruction
}

func (iLoad3 *ILoad3) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 3)
}
