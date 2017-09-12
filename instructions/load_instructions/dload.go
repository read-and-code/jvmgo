package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dload
// Load double from local variable.
// The index is an unsigned byte.
// Both index and index+1 must be indices into the local variable array of the current frame.
// The local variable at index must contain a double.
// The value of the local variable at index is pushed onto the operand stack.
type DLoad struct {
	base_instructions.Index8Instruction
}

func (dLoad *DLoad) Execute(frame *runtime_data_area.Frame) {
	loadDoubleValueAndPush(frame, uint(dLoad.Index))
}

func loadDoubleValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetDoubleValue(index)

	frame.GetOperandStack().PushDoubleValue(value)
}
