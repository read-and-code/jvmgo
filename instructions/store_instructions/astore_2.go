package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_2
// Store reference into local variable.
// The <n> must be an index into the local variable array of the current frame.
// The objectref on the top of the operand stack must be of type returnAddress or of type reference.
// It is popped from the operand stack, and the value of the local variable at <n> is set to objectref.
type AStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (aStore2 *AStore2) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 2)
}
