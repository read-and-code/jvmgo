package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload_1
type LoadPushDoubleValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushDoubleValueByIndex1 *LoadPushDoubleValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, 1)
}
