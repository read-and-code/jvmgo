package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore_2
type PopStoreIntegerValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreIntegerValueByIndex2 *PopStoreIntegerValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, 2)
}
