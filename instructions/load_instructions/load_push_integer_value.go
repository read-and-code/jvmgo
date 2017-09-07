package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iload
type LoadPushIntegerValue struct {
	base_instructions.Index8Instruction
}

func (loadPushIntegerValue *LoadPushIntegerValue) Execute(frame *runtime_data_area.Frame) {
	loadIntegerValueAndPush(frame, uint(loadPushIntegerValue.Index))
}

func loadIntegerValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetIntegerValue(index)

	frame.GetOperandStack().PushIntegerValue(value)
}
