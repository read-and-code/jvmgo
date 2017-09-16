package conversion_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// l2i
// Convert long to int
type L2i struct {
	base_instructions.NoOperandsInstruction
}

func (l2i *L2i) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()

	operandStack.PushIntegerValue(int32(longValue))
}
