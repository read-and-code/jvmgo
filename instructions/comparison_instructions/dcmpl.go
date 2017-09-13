package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dcmpl
// Compare double.
// Both value1 and value2 must be of type double.
// The values are popped from the operand stack and undergo value set conversion, resulting in value1' and value2'.
// A floating-point comparison is performed:
// If value1' is greater than value2', the int value 1 is pushed onto the operand stack.
// Otherwise, if value1' is equal to value2', the int value 0 is pushed onto the operand stack.
// Otherwise, if value1' is less than value2', the int value -1 is pushed onto the operand stack.
// Otherwise, at least one of value1' or value2' is NaN.
// The dcmpl instruction pushes the int value -1 onto the operand stack.
type DCmpl struct {
	base_instructions.NoOperandsInstruction
}

func (dCmpl *DCmpl) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue1 := operandStack.PopDoubleValue()
	doubleValue2 := operandStack.PopDoubleValue()

	if doubleValue1 > doubleValue2 {
		operandStack.PushIntegerValue(1)
	} else if doubleValue1 == doubleValue2 {
		operandStack.PushIntegerValue(0)
	} else if doubleValue1 < doubleValue2 {
		operandStack.PushIntegerValue(-1)
	} else {
		operandStack.PushIntegerValue(-1)
	}
}
