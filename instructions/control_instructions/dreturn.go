package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// dreturn
// Return double from method
type DReturn struct {
	base_instructions.NoOperandsInstruction
}

func (dReturn *DReturn) Execute(frame *runtime_data_area.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	callerFrame := thread.GetTopFrame()
	doubleValue := currentFrame.GetOperandStack().PopDoubleValue()
	callerFrame.GetOperandStack().PushDoubleValue(doubleValue)
}
