package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_1
type PopStoreReferenceValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreReferenceValueByIndex1 *PopStoreReferenceValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 1)
}