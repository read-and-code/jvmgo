package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_1
type LoadPushFloatValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushFloatValueByIndex1 *LoadPushFloatValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 1)
}
