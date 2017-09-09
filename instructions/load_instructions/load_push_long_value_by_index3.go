package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_3
type LoadPushLongValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushLongValueByIndex3 *LoadPushLongValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 3)
}
