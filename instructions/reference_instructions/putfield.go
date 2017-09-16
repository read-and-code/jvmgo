package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// putfield
// Set field in object
type PutField struct {
	base_instructions.Index16Instruction
}

func (putField *PutField) Execute(frame *runtime_data_area.Frame) {
	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetClass()
	constantPool := currentClass.GetConstantPool()
	fieldReference := constantPool.GetConstant(putField.Index).(*heap.FieldReference)
	field := fieldReference.GetResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		if currentClass != field.GetClass() || currentMethod.GetName() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	variableIndex := field.GetVariableIndex()
	operandStack := frame.GetOperandStack()

	switch field.GetDescriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		integerValue := operandStack.PopIntegerValue()
		objectReference := operandStack.PopReferenceValue()

		if objectReference == nil {
			panic("java.lang.NullPointerException")
		}

		objectReference.GetFields().SetIntegerValue(variableIndex, integerValue)
	case 'F':
		floatValue := operandStack.PopFloatValue()
		objectReference := operandStack.PopReferenceValue()

		if objectReference == nil {
			panic("java.lang.NullPointerException")
		}

		objectReference.GetFields().SetFloatValue(variableIndex, floatValue)
	case 'J':
		longValue := operandStack.PopLongValue()
		objectReference := operandStack.PopReferenceValue()

		if objectReference == nil {
			panic("java.lang.NullPointerException")
		}

		objectReference.GetFields().SetLongValue(variableIndex, longValue)
	case 'D':
		doubleValue := operandStack.PopDoubleValue()
		objectReference := operandStack.PopReferenceValue()

		if objectReference == nil {
			panic("java.lang.NullPointerException")
		}

		objectReference.GetFields().SetDoubleValue(variableIndex, doubleValue)
	case 'L', '[':
		referenceValue := operandStack.PopReferenceValue()
		objectReference := operandStack.PopReferenceValue()

		if objectReference == nil {
			panic("java.lang.NullPointerException")
		}

		objectReference.GetFields().SetReferenceValue(variableIndex, referenceValue)
	default:
		// TODO
	}
}
