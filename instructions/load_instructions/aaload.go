package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aaload
// Load reference from array
type AALoad struct {
	base_instructions.NoOperandsInstruction
}

func (aALoad *AALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	referenceArray := arrayReference.GetReferenceArray()

	if index < 0 || index >= int32(len(referenceArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushReferenceValue(referenceArray[index])
}
