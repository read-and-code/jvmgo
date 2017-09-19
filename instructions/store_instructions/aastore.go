package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aastore
// Store into reference array
type AAStore struct {
	base_instructions.NoOperandsInstruction
}

func (aAStore *AAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	referenceValue := operandStack.PopReferenceValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetReferenceArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = referenceValue
}
