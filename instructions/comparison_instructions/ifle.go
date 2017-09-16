package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ifle
// Branch if int comparison with zero succeeds
type IfLe struct {
	base_instructions.BranchInstruction
}

func (ifLe *IfLe) Execute(frame *runtime_data_area.Frame) {
	integerValue := frame.GetOperandStack().PopIntegerValue()

	if integerValue <= 0 {
		base_instructions.JumpToBranch(frame, ifLe.Offset)
	}
}
