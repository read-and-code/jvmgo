package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload_0
type LoadPushDoubleValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushDoubleValueByIndex0 *LoadPushDoubleValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, 0)
}
