package extended_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ifnonull
// Branch if reference not null.
type IfNoNull struct {
	base_instructions.BranchInstruction
}

func (ifNoNull *IfNoNull) Execute(frame *runtime_data_area.Frame) {
	referenceValue := frame.GetOperandStack().PopReferenceValue()

	if referenceValue != nil {
		base_instructions.JumpToBranch(frame, ifNoNull.Offset)
	}
}
