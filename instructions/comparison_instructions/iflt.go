package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iflt
// Branch if int comparison with zero succeeds
type IfLt struct {
	base_instructions.BranchInstruction
}

func (ifLt *IfLt) Execute(frame *runtime_data_area.Frame) {
	integerValue := frame.GetOperandStack().PopIntegerValue()

	if integerValue < 0 {
		base_instructions.JumpToBranch(frame, ifLt.Offset)
	}
}
