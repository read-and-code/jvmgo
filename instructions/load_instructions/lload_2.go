package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_2
// Load long from local variable
type LLoad2 struct {
	base_instructions.NoOperandsInstruction
}

func (lLoad2 *LLoad2) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 2)
}
