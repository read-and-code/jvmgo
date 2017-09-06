package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aconst_null
type PushNull struct {
	base.NoOperandsInstruction
}

func (pushNull *PushNull) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushReferenceValue(nil)
}
