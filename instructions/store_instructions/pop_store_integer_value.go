package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// istore
type PopStoreIntegerValue struct {
	base_instructions.Index8Instruction
}

func (popStoreIntegerValue *PopStoreIntegerValue) Execute(frame *runtime_data_area.Frame) {
	popIntegerValueAndStore(frame, uint(popStoreIntegerValue.Index))
}

func popIntegerValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopIntegerValue()
	frame.GetLocalVariables().SetIntegerValue(index, value)
}
