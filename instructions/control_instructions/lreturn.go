package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lreturn
// Return long from method
type LReturn struct {
	base_instructions.NoOperandsInstruction
}

func (lReturn *LReturn) Execute(frame *runtime_data_area.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	callerFrame := thread.GetTopFrame()
	longValue := currentFrame.GetOperandStack().PopLongValue()
	callerFrame.GetOperandStack().PushLongValue(longValue)
}
