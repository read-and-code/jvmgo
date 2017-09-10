package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore_0
type PopStoreIntegerValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreIntegerValueByIndex0 *PopStoreIntegerValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, 0)
}
