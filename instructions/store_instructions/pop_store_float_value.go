package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fstore
type PopStoreFloatValue struct {
	base_instructions.Index8Instruction
}

func (popStoreFloatValue *PopStoreFloatValue) Execute(frame *runtime_data_area.Frame) {
	popFloatValueAndStore(frame, uint(popStoreFloatValue.Index))
}

func popFloatValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopFloatValue()
	frame.GetLocalVariables().SetFloatValue(index, value)
}
