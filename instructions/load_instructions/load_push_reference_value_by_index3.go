package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_3
type LoadPushReferenceValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (loadPushReferenceValueByIndex3 *LoadPushReferenceValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 3)
}
