package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore_2
type PopStoreFloatValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreFloatValueByIndex2 *PopStoreFloatValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, 2)
}
