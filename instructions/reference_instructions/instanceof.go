package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// instanceof
// Determine if object is of given type
type InstanceOf struct {
	base_instructions.Index16Instruction
}

func (instanceOf *InstanceOf) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	objectReference := operandStack.PopReferenceValue()

	if objectReference == nil {
		operandStack.PushIntegerValue(0)

		return
	}

	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classReference := constantPool.GetConstant(instanceOf.Index).(*heap.ClassReference)
	class := classReference.GetResolvedClass()

	if objectReference.IsInstanceOf(class) {
		operandStack.PushIntegerValue(1)
	} else {
		operandStack.PushIntegerValue(0)
	}
}
