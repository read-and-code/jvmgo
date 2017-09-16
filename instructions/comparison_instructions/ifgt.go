package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ifgt
// Branch if int comparison with zero succeeds
type IfGt struct {
	base_instructions.BranchInstruction
}

func (ifGt *IfGt) Execute(frame *runtime_data_area.Frame) {
	integerValue := frame.GetOperandStack().PopIntegerValue()

	if integerValue > 0 {
		base_instructions.JumpToBranch(frame, ifGt.Offset)
	}
}
