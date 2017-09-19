package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// baload
// Load byte or boolean from array
type BALoad struct {
	base_instructions.NoOperandsInstruction
}

func (bALoad *BALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	byteArray := arrayReference.GetByteArray()

	if index < 0 || index >= int32(len(byteArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushIntegerValue(int32(byteArray[index]))
}
