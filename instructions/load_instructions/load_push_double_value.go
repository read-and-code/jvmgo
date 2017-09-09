package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload
type LoadPushDoubleValue struct {
	base_instructions.Index8Instruction
}

func (loadPushDoubleValue *LoadPushDoubleValue) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, uint(loadPushDoubleValue.Index))
}

func loadDoubleValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetDoubleValue(index)

	frame.GetOperandStack().PushDoubleValue(value)
}
