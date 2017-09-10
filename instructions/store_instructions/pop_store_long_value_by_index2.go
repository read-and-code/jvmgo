package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_2
type PopStoreLongValueByIndex2 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreLongValueByIndex2 *PopStoreLongValueByIndex2) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 2)
}
