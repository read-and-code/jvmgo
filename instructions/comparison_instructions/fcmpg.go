package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fcmpg
// Compare float
type FCmpg struct {
	base_instructions.NoOperandsInstruction
}

func (fCmpg *FCmpg) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue2 := operandStack.PopFloatValue()
	floatValue1 := operandStack.PopFloatValue()

	if floatValue1 > floatValue2 {
		operandStack.PushIntegerValue(1)
	} else if floatValue1 == floatValue2 {
		operandStack.PushIntegerValue(0)
	} else if floatValue1 < floatValue2 {
		operandStack.PushIntegerValue(-1)
	} else {
		operandStack.PushIntegerValue(1)
	}
}
