package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// saload
// Load short from array
type SALoad struct {
	base_instructions.NoOperandsInstruction
}

func (sALoad *SALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	shortArray := arrayReference.GetShortArray()

	if index < 0 || index >= int32(len(shortArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushIntegerValue(int32(shortArray[index]))
}
