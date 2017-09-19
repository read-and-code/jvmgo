package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// faload
// Load float from array
type FALoad struct {
	base_instructions.NoOperandsInstruction
}

func (fALoad *FALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	floatArray := arrayReference.GetFloatArray()

	if index < 0 || index >= int32(len(floatArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushFloatValue(floatArray[index])
}
