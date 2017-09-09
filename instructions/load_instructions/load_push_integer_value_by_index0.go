package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_0
type LoadPushIntegerValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushIntegerValueByIndex0 *LoadPushIntegerValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 0)
}
