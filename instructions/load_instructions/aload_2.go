package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload_2
// Load reference from local variable.
// The <n> must be an index into the local variable array of the current frame.
// The local variable at <n> must contain a reference.
// The objectref in the local variable at <n> is pushed onto the operand stack.
type ALoad2 struct {
	base_instructions.NoOperandsInstruction
}

func (aLoad2 *ALoad2) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, 2)
}
