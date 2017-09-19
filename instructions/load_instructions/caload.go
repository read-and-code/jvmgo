package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// caload
// Load char from array
type CALoad struct {
	base_instructions.NoOperandsInstruction
}

func (cALoad *CALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	charArray := arrayReference.GetCharArray()

	if index < 0 || index >= int32(len(charArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushIntegerValue(int32(charArray[index]))
}
