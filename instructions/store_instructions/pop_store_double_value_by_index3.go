package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore_3
type PopStoreDoubleValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreDoubleValueByIndex3 *PopStoreDoubleValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, 3)
}
