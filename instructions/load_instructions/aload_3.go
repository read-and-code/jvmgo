package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_3
// Load reference from local variable
type ALoad3 struct {
	base_instructions.NoOperandsInstruction
}

func (aLoad3 *ALoad3) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 3)
}
