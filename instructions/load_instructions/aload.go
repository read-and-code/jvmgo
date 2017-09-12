package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload
// Load reference from local variable.
// The index is an unsigned byte that must be an index into the local variable array of the current frame.
// The local variable at index must contain a reference.
// The objectref in the local variable at index is pushed onto the operand stack.
type ALoad struct {
	base_instructions.Index8Instruction
}

func (aLoad *ALoad) Execute(frame *runtime_data_area.Frame) {
	loadReferenceValueAndPush(frame, uint(aLoad.Index))
}

func loadReferenceValueAndPush(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetLocalVariables().GetReferenceValue(index)

	frame.GetOperandStack().PushReferenceValue(value)
}
