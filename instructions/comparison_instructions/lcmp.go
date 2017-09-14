package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lcmp
// Compare long.
// Both value1 and value2 must be of type long.
// They are both popped from the operand stack, and a signed integer comparison is performed.
// If value1 is greater than value2, the int value 1 is pushed onto the operand stack.
// If value1 is equal to value2, the int value 0 is pushed onto the operand stack.
// If value1 is less than value2, the int value -1 is pushed onto the operand stack.
type LCmp struct {
	base_instructions.NoOperandsInstruction
}

func (lCmp *LCmp) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue2 := operandStack.PopLongValue()
	longValue1 := operandStack.PopLongValue()

	if longValue1 > longValue2 {
		operandStack.PushIntegerValue(1)
	} else if longValue1 == longValue2 {
		operandStack.PushIntegerValue(0)
	} else {
		operandStack.PushIntegerValue(-1)
	}
}
