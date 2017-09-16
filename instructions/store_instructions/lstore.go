package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore
// Store long into local variable
type LStore struct {
	base_instructions.Index8Instruction
}

func (lStore *LStore) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, uint(lStore.Index))
}

func popLongValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopLongValue()
	frame.GetLocalVariables().SetLongValue(index, value)
}
