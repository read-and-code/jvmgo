package base_instructions

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

func InvokeMethod(frame *runtime_data_area.Frame, method *heap.Method) {
	thread := frame.GetThread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argumentsCount := int(method.GetArgumentsCount())

	if argumentsCount > 0 {
		for i := argumentsCount - 1; i >= 0; i-- {
			operand := frame.GetOperandStack().PopOperand()

			newFrame.GetLocalVariables().SetVariable(uint(i), operand)
		}
	}

	if method.IsNative() {
		if method.GetName() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("Native method: %v.%v%v\n", method.GetClass().GetName(), method.GetName(), method.GetDescriptor()))
		}
	}
}
