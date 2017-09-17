package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// invokestatic
// Invoke a class (static) method
type InvokeStatic struct {
	base_instructions.Index16Instruction
}

func (invokeStatic *InvokeStatic) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	methodReference := constantPool.GetConstant(invokeStatic.Index).(*heap.MethodReference)
	resolvedMethod := methodReference.GetResolvedMethod()

	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.GetClass()

	if !class.IsInitializationStarted() {
		frame.RevertNextPC()
		base_instructions.InitializeClass(frame.GetThread(), class)

		return
	}

	if !class.IsInitializationStarted() {
		frame.RevertNextPC()
		base_instructions.InitializeClass(frame.GetThread(), class)

		return
	}

	base_instructions.InvokeMethod(frame, resolvedMethod)
}
