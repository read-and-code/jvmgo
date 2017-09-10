package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore_0
type PopStoreLongValueByIndex0 struct {
	base_instructions.NoOperandsInstruction
}

func (popStoreLongValueByIndex0 *PopStoreLongValueByIndex0) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, 0)
}
