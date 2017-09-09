package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_0
type LoadPushLongValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushLongValueByIndex0 *LoadPushLongValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 0)
}
