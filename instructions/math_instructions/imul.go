package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// imul
// Multiply int
type IMul struct {
	base_instructions.NoOperandsInstruction
}

func (iMul *IMul) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue2 := operandStack.PopIntegerValue()
	integerValue1 := operandStack.PopIntegerValue()

	operandStack.PushIntegerValue(integerValue1 * integerValue2)
}
