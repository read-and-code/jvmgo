package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dastore
// Store into double array
type DAStore struct {
	base_instructions.NoOperandsInstruction
}

func (dAStore *DAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	doubleValue := operandStack.PopDoubleValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetDoubleArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = float64(doubleValue)
}
