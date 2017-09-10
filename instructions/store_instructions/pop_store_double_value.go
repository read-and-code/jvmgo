package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dstore
type PopStoreDoubleValue struct {
	base_instructions.Index8Instruction
}

func (popStoreDoubleValue *PopStoreDoubleValue) Execute(frame *runtime_data_area.Frame) {
	popDoubleValueAndStore(frame, uint(popStoreDoubleValue.Index))
}

func popDoubleValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopDoubleValue()
	frame.GetLocalVariables().SetDoubleValue(index, value)
}
