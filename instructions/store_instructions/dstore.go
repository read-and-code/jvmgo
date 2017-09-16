package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore
// Store double into local variable
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
