package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fmul
// Multiply float.
// Both value1 and value2 must be of type float.
// The values are popped from the operand stack and undergo value set conversion, resulting in value1' and value2'.
// The float result is value1' * value2'.
// The result is pushed onto the operand stack.
type FMul struct {
	base_instructions.NoOperandsInstruction
}

func (fMul *FMul) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue1 := operandStack.PopFloatValue()
	floatValue2 := operandStack.PopFloatValue()

	operandStack.PushFloatValue(floatValue1 * floatValue2)
}
