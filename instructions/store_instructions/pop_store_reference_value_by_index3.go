package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore_3
type PopStoreReferenceValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreReferenceValueByIndex3 *PopStoreReferenceValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, 3)
}
