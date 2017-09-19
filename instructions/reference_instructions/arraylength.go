package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// arraylength
// Get length of array
type ArrayLength struct {
	base_instructions.NoOperandsInstruction
}

func (arrayLength *ArrayLength) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	operandStack.PushIntegerValue(arrayReference.GetArrayLength())
}
