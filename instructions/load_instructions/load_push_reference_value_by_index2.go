package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_2
type LoadPushReferenceValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushReferenceValueByIndex2 *LoadPushReferenceValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 2)
}
