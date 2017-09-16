package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload_0
// Load long from local variable
type LLoad0 struct {
	base_instructions.NoOperandsInstruction
}

func (lLoad0 *LLoad0) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, 0)
}
