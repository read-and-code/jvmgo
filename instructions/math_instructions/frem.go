package math_instructions

import (
	"math"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// frem
// Remainder float
type FRem struct {
	base_instructions.NoOperandsInstruction
}

func (fRem *FRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue2 := operandStack.PopFloatValue()
	floatValue1 := operandStack.PopFloatValue()

	operandStack.PushFloatValue(float32(math.Mod(float64(floatValue1), float64(floatValue2))))
}
