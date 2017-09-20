package lang

import (
	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

const javaLangSystem = "java/lang/System"

func init() {
	native_methods.RegisterNativeMethod(javaLangSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

func arraycopy(frame *runtime_data_area.Frame) {
	localVariables := frame.GetLocalVariables()
	sourceArray := localVariables.GetReferenceValue(0)
	sourceArrayPosition := localVariables.GetIntegerValue(1)
	targetArray := localVariables.GetReferenceValue(2)
	targetArrayPosition := localVariables.GetIntegerValue(3)
	length := localVariables.GetIntegerValue(4)

	if sourceArray == nil || targetArray == nil {
		panic("java.lang.NullPointerException")
	}

	if !checkArrayCopy(sourceArray, targetArray) {
		panic("java.lang.ArrayStoreException")
	}

	if sourceArrayPosition < 0 || targetArrayPosition < 0 || length < 0 ||
		sourceArrayPosition+length > sourceArray.GetArrayLength() ||
		targetArrayPosition+length > targetArray.GetArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.CopyArray(sourceArray, targetArray, sourceArrayPosition, targetArrayPosition, length)
}

func checkArrayCopy(sourceArray, targetArray *heap.Object) bool {
	sourceArrayClass := sourceArray.GetClass()
	targetArrayClass := targetArray.GetClass()

	if !sourceArrayClass.IsArray() || !targetArrayClass.IsArray() {
		return false
	}

	if sourceArrayClass.GetArrayElementClass().IsPrimitive() ||
		targetArrayClass.GetArrayElementClass().IsPrimitive() {
		return sourceArrayClass == targetArrayClass
	}

	return true
}
