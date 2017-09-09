package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload_2
type LoadPushIntegerValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushIntegerValueByIndex2 *LoadPushIntegerValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, 2)
}
