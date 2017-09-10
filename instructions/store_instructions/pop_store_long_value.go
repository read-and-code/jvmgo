package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lstore
type PopStoreLongValue struct {
	base_instructions.Index8Instruction
}

func (popStoreLongValue *PopStoreLongValue) Execute(frame *runtime_data_area.Frame) {
	popLongValueAndStore(frame, uint(popStoreLongValue.Index))
}

func popLongValueAndStore(frame *runtime_data_area.Frame, index uint) {
	value := frame.GetOperandStack().PopLongValue()
	frame.GetLocalVariables().SetLongValue(index, value)
}
