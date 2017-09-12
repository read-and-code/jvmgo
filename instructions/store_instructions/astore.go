package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// astore
// Store reference into local variable.
// The index is an unsigned byte that must be an index into the local variable array of the current frame.
// The objectref on the top of the operand stack must be of type returnAddress or of type reference.
// It is popped from the operand stack, and the value of the local variable at index is set to objectref.
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
