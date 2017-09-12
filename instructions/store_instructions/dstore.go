package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore
// Store double into local variable.
// The index is an unsigned byte.
// Both index and index+1 must be indices into the local variable array of the current frame.
// The value on the top of the operand stack must be of type double.
// It is popped from the operand stack and undergoes value set conversion, resulting in value'.
// The local variables at index and index+1 are set to value'.
type DStore struct {
	base_instructions.Index8Instruction
}

func (dStore *DStore) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, uint(dStore.Index))
}

func popDoubleValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopDoubleValue()
	frame.GetLocalVariables().SetDoubleValue(index, value)
}
