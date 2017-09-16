package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore
// Store reference into local variable
type AStore struct {
	base_instructions.Index8Instruction
}

func (aStore *AStore) Execute(frame *runtime_data_area.Frame) {
	popReferenceValueAndStore(frame, uint(aStore.Index))
}

func popReferenceValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopReferenceValue()
	frame.GetLocalVariables().SetReferenceValue(index, value)
}
