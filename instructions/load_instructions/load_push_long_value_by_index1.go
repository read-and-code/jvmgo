package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_1
type LoadPushLongValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushLongValueByIndex1 *LoadPushLongValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 1)
}
