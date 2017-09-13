package math_instructions

import (
	"math"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// drem
// Remainder double.
// Both value1 and value2 must be of type double.
// The values are popped from the operand stack and undergo value set conversion, resulting in value1' and value2'.
// The result is calculated and pushed onto the operand stack as a double.
type DRem struct {
	base_instructions.NoOperandsInstruction
}

func (dRem *DRem) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue1 := operandStack.PopDoubleValue()
	doubleValue2 := operandStack.PopDoubleValue()

	operandStack.PushDoubleValue(math.Mod(doubleValue1, doubleValue2))
}
