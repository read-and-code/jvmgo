package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// getfield
// Fetch field from object
type GetField struct {
	base_instructions.Index16Instruction
}

func (getField *GetField) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	fieldReference := constantPool.GetConstant(getField.Index).(*heap.FieldReference)
	field := fieldReference.GetResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	operandStack := frame.GetOperandStack()
	objectReference := operandStack.PopReferenceValue()

	if objectReference == nil {
		panic("java.lang.NullPointerException")
	}

	variableIndex := field.GetVariableIndex()
	fields := objectReference.GetFields()

	switch field.GetDescriptor()[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		operandStack.PushIntegerValue(fields.GetIntegerValue(variableIndex))
	case 'F':
		operandStack.PushFloatValue(fields.GetFloatValue(variableIndex))
	case 'J':
		operandStack.PushLongValue(fields.GetLongValue(variableIndex))
	case 'D':
		operandStack.PushDoubleValue(fields.GetDoubleValue(variableIndex))
	case 'L', '[':
		operandStack.PushReferenceValue(fields.GetReferenceValue(variableIndex))
	default:
		// TODO
	}
}
