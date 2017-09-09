package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_2
type LoadPushFloatValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushFloatValueByIndex2 *LoadPushFloatValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 2)
}
