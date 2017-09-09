package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_0
type LoadPushReferenceValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushReferenceValueByIndex0 *LoadPushReferenceValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 0)
}
