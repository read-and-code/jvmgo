package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore_0
type PopStoreFloatValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreFloatValueByIndex0 *PopStoreFloatValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, 0)
}
