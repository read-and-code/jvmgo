package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_2
type PopStoreReferenceValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreReferenceValueByIndex2 *PopStoreReferenceValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 2)
}
