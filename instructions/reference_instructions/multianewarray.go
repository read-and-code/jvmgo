package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// multianewarray
// Create new multidimensional array
type MultiANewArray struct {
	index      uint16
	dimensions uint8
}

func (multiANewArray *MultiANewArray) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	multiANewArray.index = bytecodeReader.ReadUint16()
	multiANewArray.dimensions = bytecodeReader.ReadUint8()
}

func (multiANewArray *MultiANewArray) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classReference := constantPool.GetConstant(uint(multiANewArray.index)).(*heap.ClassReference)
	arrayClass := classReference.GetResolvedClass()

	operandStack := frame.GetOperandStack()
	arrayLengthInEachDimension := getAndCheckArrayLengthInEachDimension(operandStack, int(multiANewArray.dimensions))
	array := newMultidimensionalArray(arrayLengthInEachDimension, arrayClass)

	operandStack.PushReferenceValue(array)
}

func getAndCheckArrayLengthInEachDimension(operandStack *runtime_data_area.OperandStack, dimensions int) []int32 {
	arrayLengthInEachDimension := make([]int32, dimensions)

	for i := dimensions - 1; i >= 0; i-- {
		arrayLengthInEachDimension[i] = operandStack.PopIntegerValue()

		if arrayLengthInEachDimension[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return arrayLengthInEachDimension
}

func newMultidimensionalArray(arrayLengthInEachDimension []int32, arrayClass *heap.Class) *heap.Object {
	arrayLength := uint(arrayLengthInEachDimension[0])
	array := arrayClass.NewArray(arrayLength)

	if len(arrayLengthInEachDimension) > 1 {
		referenceArray := array.GetReferenceArray()

		for i := range referenceArray {
			referenceArray[i] = newMultidimensionalArray(arrayLengthInEachDimension[1:], arrayClass.GetArrayElementClass())
		}
	}

	return array
}
