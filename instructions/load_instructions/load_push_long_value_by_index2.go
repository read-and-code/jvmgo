package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_2
type LoadPushLongValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushLongValueByIndex2 *LoadPushLongValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 2)
}
