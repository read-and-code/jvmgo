package reserved_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/native_methods"
	_ "github.com/Frederick-S/jvmgo/native_methods/java/lang"
	_ "github.com/Frederick-S/jvmgo/native_methods/sun/misc"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

type InvokeNative struct {
	base_instructions.NoOperandsInstruction
}

func (invokeNative *InvokeNative) Execute(frame *runtime_data_area.Frame) {
	method := frame.GetMethod()
	className := method.GetClass().GetName()
	methodName := method.GetName()
	methodDescriptor := method.GetDescriptor()

	nativeMethod := native_methods.FindNativeMethod(className, methodName, methodDescriptor)

	if nativeMethod == nil {
		panic("java.lang.UnsatisfiedLinkError: " + className + "." + methodName + methodDescriptor)
	}

	nativeMethod(frame)
}
