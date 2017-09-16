package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload
// Load long from local variable
type LLoad struct {
	base_instructions.Index8Instruction
}

func (lLoad *LLoad) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, uint(lLoad.Index))
}

func loadLongValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetLongValue(index)

	frame.GetOperandStack().PushLongValue(value)
}
