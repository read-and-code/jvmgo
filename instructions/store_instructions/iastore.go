package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iastore
// Store into int array
type IAStore struct {
	base_instructions.NoOperandsInstruction
}

func (iAStore *IAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue := operandStack.PopIntegerValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetIntArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = int32(integerValue)
}
