package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iaload
// Load int from array
type IALoad struct {
	base_instructions.NoOperandsInstruction
}

func (iALoad *IALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	intArray := arrayReference.GetIntArray()

	if index < 0 || index >= int32(len(intArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushIntegerValue(intArray[index])
}
