package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore_1
type PopStoreFloatValueByIndex1 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreFloatValueByIndex1 *PopStoreFloatValueByIndex1) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, 1)
}
