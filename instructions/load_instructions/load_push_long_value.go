package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lload
type LoadPushLongValue struct {
	base_instructions.Index8Instruction
}

func (loadPushLongValue *LoadPushLongValue) Execute(frame *runtime_data_area.Frame) {
	loadLongValueAndPush(frame, uint(loadPushLongValue.Index))
}

func loadLongValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetLongValue(index)

	frame.GetOperandStack().PushLongValue(value)
}
