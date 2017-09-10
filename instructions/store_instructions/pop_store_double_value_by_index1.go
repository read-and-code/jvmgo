package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_1
type PopStoreDoubleValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreDoubleValueByIndex1 *PopStoreDoubleValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 1)
}
