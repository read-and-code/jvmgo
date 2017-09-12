package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fload
// Load float from local variable.
// The index is an unsigned byte that must be an index into the local variable array of the current frame.
// The local variable at index must contain a float.
// The value of the local variable at index is pushed onto the operand stack.
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
