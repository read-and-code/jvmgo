package store_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// castore
// Store into char array
type CAStore struct {
	base_instructions.NoOperandsInstruction
}

func (cAStore *CAStore) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	charValue := operandStack.PopIntegerValue()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetCharArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	referenceArray[index] = uint16(charValue)
}
