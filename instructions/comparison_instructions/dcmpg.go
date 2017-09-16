package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dcmpg
// Compare double
type DCmpg struct {
	base_instructions.NoOperandsInstruction
}

func (dCmpg *DCmpg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue2 := operandStack.PopDoubleValue()
	doubleValue1 := operandStack.PopDoubleValue()

	if doubleValue1 > doubleValue2 {
		operandStack.PushIntegerValue(1)
	} else if doubleValue1 == doubleValue2 {
		operandStack.PushIntegerValue(0)
	} else if doubleValue1 < doubleValue2 {
		operandStack.PushIntegerValue(-1)
	} else {
		operandStack.PushIntegerValue(1)
	}
}
