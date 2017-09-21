package reference_instructions

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// invokevirtual
// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	base_instructions.Index16Instruction
}

func (invokeVirtual *InvokeVirtual) Execute(frame *runtime_data_area.Frame) {
	currentClass := frame.GetMethod().GetClass()
	constantPool := currentClass.GetConstantPool()
	methodReference := constantPool.GetConstant(invokeVirtual.Index).(*heap.MethodReference)
	resolvedMethod := methodReference.GetResolvedMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	fmt.Println("-----------------------------" + resolvedMethod.GetName())

	referenceValue := frame.GetOperandStack().GetReferenceValueBelowTop(resolvedMethod.GetArgumentsCount() - 1)

	if referenceValue == nil {
		if methodReference.GetName() == "println" {
			printLine(frame.GetOperandStack(), methodReference.GetDescriptor())

			return
		}

		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() && resolvedMethod.GetClass().IsSuperClassOf(currentClass) &&
		resolvedMethod.GetClass().GetPackageName() != currentClass.GetPackageName() &&
		referenceValue.GetClass() != currentClass &&
		!referenceValue.GetClass().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(referenceValue.GetClass(), methodReference.GetName(), methodReference.GetDescriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base_instructions.InvokeMethod(frame, methodToBeInvoked)
}

func printLine(operandStack *runtime_data_area.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", operandStack.PopIntegerValue() != 0)
	case "(C)V":
		fmt.Printf("%c\n", operandStack.PopIntegerValue())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", operandStack.PopIntegerValue())
	case "(F)V":
		fmt.Printf("%v\n", operandStack.PopFloatValue())
	case "(J)V":
		fmt.Printf("%v\n", operandStack.PopLongValue())
	case "(D)V":
		fmt.Printf("%v\n", operandStack.PopDoubleValue())
	case "(Ljava/lang/String;)V":
		javaString := operandStack.PopReferenceValue()
		goString := heap.ConvertJavaStringToGoString(javaString)

		fmt.Println(goString)
	default:
		panic("println: " + descriptor)
	}

	operandStack.PopReferenceValue()
}
