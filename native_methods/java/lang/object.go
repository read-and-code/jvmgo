package lang

import (
	"unsafe"

	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

const javaLangObject = "java/lang/Object"

func init() {
	native_methods.RegisterNativeMethod(javaLangObject, "getClass", "()Ljava/lang/Class;", getClass)
	native_methods.RegisterNativeMethod(javaLangObject, "hashCode", "()I;", getHashCode)
	native_methods.RegisterNativeMethod(javaLangObject, "clone", "()Ljava/lang/Object;", clone)
}

func getClass(frame *runtime_data_area.Frame) {
	this := frame.GetLocalVariables().GetThis()
	class := this.GetClass().GetJavaClass()

	frame.GetOperandStack().PushReferenceValue(class)
}

func getHashCode(frame *runtime_data_area.Frame) {
	this := frame.GetLocalVariables().GetThis()
	hashCode := int32(uintptr(unsafe.Pointer(this)))

	frame.GetOperandStack().PushIntegerValue(hashCode)
}

func clone(frame *runtime_data_area.Frame) {
	this := frame.GetLocalVariables().GetThis()

	cloneable := frame.GetLocalVariables().GetThis().GetClass().GetClassLoader().LoadClass("java/lang/Cloneable")

	if !this.GetClass().IsImplementsFrom(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.GetOperandStack().PushReferenceValue(this.Clone())
}
