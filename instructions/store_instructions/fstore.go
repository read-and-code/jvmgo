package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore
// Store float into local variable.
// The index is an unsigned byte that must be an index into the local variable array of the current frame.
// The value on the top of the operand stack must be of type float.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// The value of the local variable at index is set to value'.
type FStore struct {
	base_instructions.Index8Instruction
}

func (fStore *FStore) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, uint(fStore.Index))
}

func popFloatValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopFloatValue()
	frame.GetLocalVariables().SetFloatValue(index, value)
}
