package extended_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ifnull
// Branch if reference is null
type IfNull struct {
	base_instructions.BranchInstruction
}

func (ifNull *IfNull) Execute(frame *runtime_data_area.Frame) {
	referenceValue := frame.GetOperandStack().PopReferenceValue()

	if referenceValue == nil {
		base_instructions.JumpToBranch(frame, ifNull.Offset)
	}
}
