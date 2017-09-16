package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_0
// Load reference from local variable
type ALoad0 struct {
	base_instructions.NoOperandsInstruction
}

func (aLoad0 *ALoad0) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 0)
}
