package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lcmp
// Compare long
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
