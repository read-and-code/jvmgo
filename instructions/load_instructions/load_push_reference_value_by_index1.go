package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_1
type LoadPushReferenceValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushReferenceValueByIndex1 *LoadPushReferenceValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 1)
}
