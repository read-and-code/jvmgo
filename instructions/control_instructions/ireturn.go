package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ireturn
// Return int from method
type IReturn struct {
	base_instructions.NoOperandsInstruction
}

func (iReturn *IReturn) Execute(frame *runtime_data_area.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	callerFrame := thread.GetTopFrame()
	integerValue := currentFrame.GetOperandStack().PopIntegerValue()
	callerFrame.GetOperandStack().PushIntegerValue(integerValue)
}
