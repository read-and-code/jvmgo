package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// invokespecial
// Invoke instance method; special handling for superclass, private, and instance initialization method invocations
type InvokeSpecial struct {
	base_instructions.Index16Instruction
}

func (invokeSpecial *InvokeSpecial) Execute(frame *runtime_data_area.Frame) {
	currentClass := frame.GetMethod().GetClass()
	constantPool := currentClass.GetConstantPool()
	methodReference := constantPool.GetConstant(invokeSpecial.Index).(*heap.MethodReference)
	resolvedClass := methodReference.GetResolvedClass()
	resolvedMethod := methodReference.GetResolvedMethod()

	if resolvedMethod.GetName() != "<init>" && resolvedMethod.GetClass() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	referenceValue := frame.GetOperandStack().GetReferenceValueBelowTop(resolvedMethod.GetArgumentsCount() - 1)

	if referenceValue == nil {
		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() && resolvedMethod.GetClass().IsSuperClassOf(currentClass) &&
		resolvedMethod.GetClass().GetPackageName() != currentClass.GetPackageName() &&
		referenceValue.GetClass() != currentClass &&
		!referenceValue.GetClass().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod

	if currentClass.IsSuper() && resolvedClass.IsSuperClassOf(currentClass) && resolvedMethod.GetName() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.GetSuperClass(), methodReference.GetName(), methodReference.GetDescriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base_instructions.InvokeMethod(frame, methodToBeInvoked)
}
