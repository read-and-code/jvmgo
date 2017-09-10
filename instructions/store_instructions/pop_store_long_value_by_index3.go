package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_3
type PopStoreLongValueByIndex3 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreLongValueByIndex3 *PopStoreLongValueByIndex3) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 3)
}
