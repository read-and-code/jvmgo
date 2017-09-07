package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_3
type LoadPushIntegerValueByIndex3 struct {
	base_instructions.Index8Instruction
}

func (loadPushIntegerValueByIndex3 *LoadPushIntegerValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 3)
}
