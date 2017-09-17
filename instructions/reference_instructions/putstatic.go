package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// putstatic
// Set static field in class
type PutStatic struct {
	base_instructions.Index16Instruction
}

func (putStatic *PutStatic) Execute(frame *runtime_data_area.Frame) {
	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetClass()
	constantPool := currentClass.GetConstantPool()
	fieldReference := constantPool.GetConstant(putStatic.Index).(*heap.FieldReference)
	field := fieldReference.GetResolvedField()
	class := field.GetClass()

	if !class.IsInitializationStarted() {
		frame.RevertNextPC()
		base_instructions.InitializeClass(frame.GetThread(), class)

		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		if currentClass != class || currentMethod.GetName() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	variableIndex := field.GetVariableIndex()
	staticVariables := class.GetStaticVariables()
	operandStack := frame.GetOperandStack()

	switch field.GetDescriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		staticVariables.SetIntegerValue(variableIndex, operandStack.PopIntegerValue())
	case 'F':
		staticVariables.SetFloatValue(variableIndex, operandStack.PopFloatValue())
	case 'J':
		staticVariables.SetLongValue(variableIndex, operandStack.PopLongValue())
	case 'D':
		staticVariables.SetDoubleValue(variableIndex, operandStack.PopDoubleValue())
	case 'L', '[':
		staticVariables.SetReferenceValue(variableIndex, operandStack.PopReferenceValue())
	default:
		// TODO
	}
}
