package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload
// Load float from local variable
type FLoad struct {
	base_instructions.Index8Instruction
}

func (fLoad *FLoad) Execute(frame *runtime_data_area.Frame) {
	loadFloatValueAndPush(frame, uint(fLoad.Index))
}

func loadFloatValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetFloatValue(index)

	frame.GetOperandStack().PushFloatValue(value)
}
