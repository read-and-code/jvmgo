package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// checkcast
// Check whether object is of given type
type CheckCast struct {
	base_instructions.Index16Instruction
}

func (checkCast *CheckCast) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	objectReference := operandStack.PopReferenceValue()

	operandStack.PushReferenceValue(objectReference)

	if objectReference == nil {
		return
	}

	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classReference := constantPool.GetConstant(checkCast.Index).(*heap.ClassReference)
	class := classReference.GetResolvedClass()

	if !objectReference.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
