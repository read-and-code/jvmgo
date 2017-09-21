package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// getstatic
// Get static field from class
type GetStatic struct {
	base_instructions.Index16Instruction
}

func (getStatic *GetStatic) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	fieldReference := constantPool.GetConstant(getStatic.Index).(*heap.FieldReference)
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

	variableIndex := field.GetVariableIndex()
	staticVariables := class.GetStaticVariables()
	operandStack := frame.GetOperandStack()

	switch field.GetDescriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		operandStack.PushIntegerValue(staticVariables.GetIntegerValue(variableIndex))
	case 'F':
		operandStack.PushFloatValue(staticVariables.GetFloatValue(variableIndex))
	case 'J':
		operandStack.PushLongValue(staticVariables.GetLongValue(variableIndex))
	case 'D':
		operandStack.PushDoubleValue(staticVariables.GetDoubleValue(variableIndex))
	case 'L', '[':
		operandStack.PushReferenceValue(staticVariables.GetReferenceValue(variableIndex))
	default:
		// TODO
	}
}
