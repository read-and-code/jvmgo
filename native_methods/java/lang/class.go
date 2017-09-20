package lang

import (
	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

const javaLangClass = "java/lang/Class"

func init() {
	native_methods.RegisterNativeMethod(javaLangClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native_methods.RegisterNativeMethod(javaLangClass, "getName0", "()Ljava/lang/String;", getName0)
	native_methods.RegisterNativeMethod(javaLangClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

func getPrimitiveClass(frame *runtime_data_area.Frame) {
	nameObject := frame.GetLocalVariables().GetReferenceValue(0)
	name := heap.ConvertJavaStringToGoString(nameObject)

	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	class := classLoader.LoadClass(name).GetJavaClass()

	frame.GetOperandStack().PushReferenceValue(class)
}

func getName0(frame *runtime_data_area.Frame) {
	class := frame.GetLocalVariables().GetThis().GetExtraData().(*heap.Class)

	name := class.GetJavaName()
	nameObject := heap.ConvertGoStringToJavaString(class.GetClassLoader(), name)

	frame.GetOperandStack().PushReferenceValue(nameObject)
}

func desiredAssertionStatus0(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushBooleanValue(false)
}

func isInterface(frame *runtime_data_area.Frame) {
	class := frame.GetLocalVariables().GetThis().GetExtraData().(*heap.Class)

	frame.GetOperandStack().PushBooleanValue(class.IsInterface())
}

func isPrimitive(frame *runtime_data_area.Frame) {
	class := frame.GetLocalVariables().GetThis().GetExtraData().(*heap.Class)

	frame.GetOperandStack().PushBooleanValue(class.IsPrimitive())
}
