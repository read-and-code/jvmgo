package math_instructions

import (
	"math"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// drem
// Remainder double
type DRem struct {
	base_instructions.NoOperandsInstruction
}

func (dRem *DRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue2 := operandStack.PopDoubleValue()
	doubleValue1 := operandStack.PopDoubleValue()

	operandStack.PushDoubleValue(math.Mod(doubleValue1, doubleValue2))
}
