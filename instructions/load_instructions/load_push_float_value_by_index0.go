package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_0
type LoadPushFloatValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushFloatValueByIndex0 *LoadPushFloatValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 0)
}
