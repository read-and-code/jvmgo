package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload_2
type LoadPushDoubleValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushDoubleValueByIndex2 *LoadPushDoubleValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, 2)
}
