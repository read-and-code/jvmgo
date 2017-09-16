package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_3
// Load long from local variable
type LLoad3 struct {
	base_instructions.NoOperandsInstruction
}

func (lLoad3 *LLoad3) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 3)
}
