package load_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// daload
// Load double from array
type DALoad struct {
	base_instructions.NoOperandsInstruction
}

func (dALoad *DALoad) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	index := operandStack.PopIntegerValue()
	arrayReference := operandStack.PopReferenceValue()

	if arrayReference == nil {
		panic("java.lang.NullPointerException")
	}

	doubleArray := arrayReference.GetDoubleArray()

	if index < 0 || index >= int32(len(doubleArray)) {
		panic("ArrayIndexOutOfBoundsException")
	}

	operandStack.PushDoubleValue(doubleArray[index])
}
