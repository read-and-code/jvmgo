package reference_instructions

import (
	"reflect"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// athrow
// Throw exception or error
type AThrow struct {
	base_instructions.NoOperandsInstruction
}

func (aThrow *AThrow) Execute(frame *runtime_data_area.Frame) {
	exception := frame.GetOperandStack().PopReferenceValue()

	if exception == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.GetThread()

	if !findAndGotoExceptionHandler(thread, exception) {
		handleUncaughtException(thread, exception)
	}
}

func findAndGotoExceptionHandler(thread *runtime_data_area.Thread, exception *heap.Object) bool {
	for {
		frame := thread.GetCurrentFrame()
		pc := frame.GetNextPC() - 1

		handlerPC := frame.GetMethod().FindExceptionHandler(exception.GetClass(), pc)

		if handlerPC > 0 {
			operandStack := frame.GetOperandStack()
			operandStack.Clear()
			operandStack.PushReferenceValue(exception)

			frame.SetNextPC(handlerPC)

			return true
		}

		thread.PopFrame()

		if thread.IsJVMStackEmpty() {
			break
		}
	}

	return false
}

func handleUncaughtException(thread *runtime_data_area.Thread, exception *heap.Object) {
	thread.ClearStack()

	javaMessage := exception.GetReferenceValue("detailMessage", "Ljava/lang/String;")
	goMessage := heap.ConvertJavaStringToGoString(javaMessage)

	println(exception.GetClass().GetJavaName() + ": " + goMessage)

	stackTraceElements := reflect.ValueOf(exception.GetExtraData())

	for i := 0; i < stackTraceElements.Len(); i++ {
		stackTraceElement := stackTraceElements.Index(i).Interface().(interface {
			String() string
		})

		println("\tat " + stackTraceElement.String())
	}
}
