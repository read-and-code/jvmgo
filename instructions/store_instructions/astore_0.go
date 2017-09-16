package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_0
// Store reference into local variable
type AStore0 struct {
	base_instructions.NoOperandsInstruction
}

func (aStore0 *AStore0) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 0)
}
