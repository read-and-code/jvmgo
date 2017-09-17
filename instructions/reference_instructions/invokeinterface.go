package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// invokeinterface
// Invoke interface method
type InvokeInterface struct {
	index uint
}

func (invokeInterface *InvokeInterface) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	invokeInterface.index = uint(bytecodeReader.ReadUint16())

	bytecodeReader.ReadUint8()
	bytecodeReader.ReadUint8()
}

func (invokeInterface *InvokeInterface) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	methodReference := constantPool.GetConstant(invokeInterface.index).(*heap.InterfaceMethodReference)
	resolvedMethod := methodReference.GetResolvedInterfaceMethod()

	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	referenceValue := frame.GetOperandStack().GetReferenceValueBelowTop(resolvedMethod.GetArgumentsCount() - 1)

	if referenceValue == nil {
		panic("java.lang.NullPointerException")
	}

	if !referenceValue.GetClass().IsImplementsFrom(methodReference.GetResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(referenceValue.GetClass(), methodReference.GetName(), methodReference.GetDescriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base_instructions.InvokeMethod(frame, methodToBeInvoked)
}
