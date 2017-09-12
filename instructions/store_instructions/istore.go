package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore
// Store int into local variable.
// The index is an unsigned byte that must be an index into the local variable array of the current frame.
// The value on the top of the operand stack must be of type int.
// It is popped from the operand stack, and the value of the local variable at index is set to value.
type IStore struct {
	base_instructions.Index8Instruction
}

func (iStore *IStore) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, uint(iStore.Index))
}

func popIntegerValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopIntegerValue()
	frame.GetLocalVariables().SetIntegerValue(index, value)
}
