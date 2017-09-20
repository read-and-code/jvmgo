package lang

import (
	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

const javaLangString = "java/lang/String"

func init() {
	native_methods.RegisterNativeMethod(javaLangString, "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *runtime_data_area.Frame) {
	this := frame.GetLocalVariables().GetThis()
	internedString := heap.GetInternedString(this)

	frame.GetOperandStack().PushReferenceValue(internedString)
}
