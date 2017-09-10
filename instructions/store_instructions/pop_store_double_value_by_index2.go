package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_2
type PopStoreDoubleValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreDoubleValueByIndex2 *PopStoreDoubleValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 2)
}
