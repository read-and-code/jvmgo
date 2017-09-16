package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_2
// Store reference into local variable
type AStore2 struct {
	base_instructions.NoOperandsInstruction
}

func (aStore2 *AStore2) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 2)
}
