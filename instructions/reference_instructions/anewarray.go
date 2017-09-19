package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// anewarray
// Create new array of reference
type ANewArray struct {
	base_instructions.Index16Instruction
}

func (aNewArray *ANewArray) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classReference := constantPool.GetConstant(aNewArray.Index).(*heap.ClassReference)
	arrayElementClass := classReference.GetResolvedClass()

	operandStack := frame.GetOperandStack()
	arrayLength := operandStack.PopIntegerValue()

	if arrayLength < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrayClass := arrayElementClass.GetArrayClass()
	array := arrayClass.NewArray(uint(arrayLength))

	operandStack.PushReferenceValue(array)
}
