package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dadd
// Add double.
// Both value1 and value2 must be of type double.
// The values are popped from the operand stack and undergo value set conversion, resulting in value1' and value2'.
// The double result is value1' + value2'. The result is pushed onto the operand stack.
type DAdd struct {
	base_instructions.NoOperandsInstruction
}

func (dAdd *DAdd) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue1 := operandStack.PopDoubleValue()
	doubleValue2 := operandStack.PopDoubleValue()

	operandStack.PushDoubleValue(doubleValue1 + doubleValue2)
}
