package lang

import (
	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

const javaLangThrowable = "java/lang/Throwable"

func init() {
	native_methods.RegisterNativeMethod(javaLangThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *runtime_data_area.Frame) {
	this := frame.GetLocalVariables().GetThis()
	frame.GetOperandStack().PushReferenceValue(this)

	stackTraceElements := createStackTraceElements(this, frame.GetThread())

	this.SetExtraData(stackTraceElements)
}

func createStackTraceElements(object *heap.Object, thread *runtime_data_area.Thread) []*StackTraceElement {
	skip := distanceToObject(object.GetClass()) + 2
	frames := thread.GetFrames()[skip:]
	stackTraceElements := make([]*StackTraceElement, len(frames))

	for i, frame := range frames {
		stackTraceElements[i] = createStackTraceElement(frame)
	}

	return stackTraceElements
}

func distanceToObject(class *heap.Class) int {
	distance := 0

	for currentClass := class.GetSuperClass(); currentClass != nil; currentClass = currentClass.GetSuperClass() {
		distance++
	}

	return distance
}

func createStackTraceElement(frame *runtime_data_area.Frame) *StackTraceElement {
	method := frame.GetMethod()
	class := method.GetClass()

	return &StackTraceElement{
		fileName:   class.GetSourceFileName(),
		className:  class.GetJavaName(),
		methodName: method.GetName(),
		lineNumber: method.GetLineNumber(frame.GetNextPC() - 1),
	}
}
