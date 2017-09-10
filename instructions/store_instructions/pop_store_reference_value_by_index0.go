package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_0
type PopStoreReferenceValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreReferenceValueByIndex0 *PopStoreReferenceValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 0)
}
