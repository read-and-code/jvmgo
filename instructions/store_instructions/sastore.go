package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// sastore
// Store into short array
type SAStore struct {
	base_instructions.NoOperandsInstruction
}

func (sAStore *SAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	shortValue := operandStack.PopIntegerValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetShortArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = int16(shortValue)
}
