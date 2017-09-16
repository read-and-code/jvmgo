package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_2
// Load reference from local variable
type ALoad2 struct {
	base_instructions.NoOperandsInstruction
}

func (aLoad2 *ALoad2) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 2)
}
