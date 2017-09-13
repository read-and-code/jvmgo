package math_instructions

import (
	"math"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// frem
// Remainder float.
// Both value1 and value2 must be of type float.
// The values are popped from the operand stack and undergo value set conversion, resulting in value1' and value2'.
// The result is calculated and pushed onto the operand stack as a float.
type FRem struct {
	base_instructions.NoOperandsInstruction
}

func (fRem *FRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue1 := operandStack.PopFloatValue()
	floatValue2 := operandStack.PopFloatValue()

	operandStack.PushFloatValue(float32(math.Mod(float64(floatValue1), float64(floatValue2))))
}
