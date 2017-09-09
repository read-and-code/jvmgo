package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_1
type LoadPushIntegerValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushIntegerValueByIndex1 *LoadPushIntegerValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 1)
}
