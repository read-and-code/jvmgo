package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lastore
// Store into long array
type LAStore struct {
	base_instructions.NoOperandsInstruction
}

func (lAStore *LAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	longValue := operandStack.PopLongValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetLongArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = int64(longValue)
}
