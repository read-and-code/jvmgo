package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ifne
// Branch if int comparison with zero succeeds
type IfNe struct {
	base_instructions.BranchInstruction
}

func (ifNe *IfNe) Execute(frame *runtime_data_area.Frame) {
	integerValue := frame.GetOperandStack().PopIntegerValue()

	if integerValue != 0 {
		base_instructions.JumpToBranch(frame, ifNe.Offset)
	}
}
