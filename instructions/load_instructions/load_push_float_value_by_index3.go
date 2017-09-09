package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload_3
type LoadPushFloatValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushFloatValueByIndex3 *LoadPushFloatValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, 3)
}
