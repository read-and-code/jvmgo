package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

const (
	ArrayTypeBoolean = 4
	ArrayTypeChar    = 5
	ArrayTypeFloat   = 6
	ArrayTypeDouble  = 7
	ArrayTypeByte    = 8
	ArrayTypeShort   = 9
	ArrayTypeInt     = 10
	ArrayTypeLong    = 11
)

// newarray
// Create new array
type NewArray struct {
	arrayType uint8
}

func (newArray *NewArray) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	newArray.arrayType = bytecodeReader.ReadUint8()
}

func (newArray *NewArray) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	arrayLength := operandStack.PopIntegerValue()

	if arrayLength < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	primitiveTypeArrayClass := getPrimitiveTypeArrayClass(classLoader, newArray.arrayType)
	array := primitiveTypeArrayClass.NewArray(uint(arrayLength))

	operandStack.PushReferenceValue(array)
}

func getPrimitiveTypeArrayClass(classLoader *heap.ClassLoader, arrayType uint8) *heap.Class {
	switch arrayType {
	case ArrayTypeBoolean:
		return classLoader.LoadClass("[Z")
	case ArrayTypeByte:
		return classLoader.LoadClass("[B")
	case ArrayTypeChar:
		return classLoader.LoadClass("[C")
	case ArrayTypeShort:
		return classLoader.LoadClass("[S")
	case ArrayTypeInt:
		return classLoader.LoadClass("[I")
	case ArrayTypeLong:
		return classLoader.LoadClass("[J")
	case ArrayTypeFloat:
		return classLoader.LoadClass("[F")
	case ArrayTypeDouble:
		return classLoader.LoadClass("[D")
	default:
		panic("Invalid array type!")
	}
}
