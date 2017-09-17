package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// areturn
// Return reference from method
type AReturn struct {
	base_instructions.NoOperandsInstruction
}

func (aReturn *AReturn) Execute(frame *runtime_data_area.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	callerFrame := thread.GetTopFrame()
	referenceValue := currentFrame.GetOperandStack().PopReferenceValue()
	callerFrame.GetOperandStack().PushReferenceValue(referenceValue)
}
