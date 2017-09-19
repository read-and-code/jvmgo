package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// laload
// Load long from array
type LALoad struct {
	base_instructions.NoOperandsInstruction
}

func (lALoad *LALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	longArray := arrayReference.GetLongArray()

	if index < 0 || index >= int32(len(longArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushLongValue(longArray[index])
}
