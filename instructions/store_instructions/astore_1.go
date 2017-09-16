package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_1
// Store reference into local variable
type AStore1 struct {
	base_instructions.NoOperandsInstruction
}

func (aStore1 *AStore1) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 1)
}
