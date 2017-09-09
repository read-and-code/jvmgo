package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload_3
type LoadPushDoubleValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushDoubleValueByIndex3 *LoadPushDoubleValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, 3)
}
