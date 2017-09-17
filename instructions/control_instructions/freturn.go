package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// freturn
// Return float from method
type FReturn struct {
	base_instructions.NoOperandsInstruction
}

func (fReturn *FReturn) Execute(frame *runtime_data_area.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	callerFrame := thread.GetTopFrame()
	floatValue := currentFrame.GetOperandStack().PopFloatValue()
	callerFrame.GetOperandStack().PushFloatValue(floatValue)
}
