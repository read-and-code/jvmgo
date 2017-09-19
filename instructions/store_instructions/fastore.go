package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// fastore
// Store into float array
type FAStore struct {
	base_instructions.NoOperandsInstruction
}

func (fAStore *FAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	floatValue := operandStack.PopFloatValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetFloatArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = float32(floatValue)
}
