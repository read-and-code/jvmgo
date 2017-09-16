package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore
// Store int into local variable
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
