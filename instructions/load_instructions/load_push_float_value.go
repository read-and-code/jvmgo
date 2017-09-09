package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload
type LoadPushFloatValue struct {
	base_instructions.Index8Instruction
}

func (loadPushFloatValue *LoadPushFloatValue) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, uint(loadPushFloatValue.Index))
}

func loadFloatValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetFloatValue(index)

	frame.GetOperandStack().PushFloatValue(value)
}
