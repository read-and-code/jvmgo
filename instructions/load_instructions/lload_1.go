package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_1
// Load long from local variable
type LLoad1 struct {
	base_instructions.NoOperandsInstruction
}

func (lLoad1 *LLoad1) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 1)
}
