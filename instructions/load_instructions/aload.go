package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aload
// Load reference from local variable
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
