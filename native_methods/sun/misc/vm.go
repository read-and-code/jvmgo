package misc

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

func init() {
	native_methods.RegisterNativeMethod("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *runtime_data_area.Frame) {
	vmClass := frame.GetMethod().GetClass()
	savedProperties := vmClass.GetReferenceVariable("savedProps", "Ljava/util/Properties;")
	key := heap.ConvertGoStringToJavaString(vmClass.GetClassLoader(), "foo")
	value := heap.ConvertGoStringToJavaString(vmClass.GetClassLoader(), "bar")

	frame.GetOperandStack().PushReferenceValue(savedProperties)
	frame.GetOperandStack().PushReferenceValue(key)
	frame.GetOperandStack().PushReferenceValue(value)

	propertiesClass := vmClass.GetClassLoader().LoadClass("java/util/Properties")
	setPropertyMethod := propertiesClass.GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")

	base_instructions.InvokeMethod(frame, setPropertyMethod)
}
