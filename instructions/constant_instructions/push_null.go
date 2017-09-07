package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aconst_null
type PushNull struct {
	base_instructions.NoOperandsInstruction
}

func (pushNull *PushNull) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushReferenceValue(nil)
}
