package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// bastore
// Store into byte or boolean array
type BAStore struct {
	base_instructions.NoOperandsInstruction
}

func (bAStore *BAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	byteValue := operandStack.PopIntegerValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetByteArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = int8(byteValue)
}
