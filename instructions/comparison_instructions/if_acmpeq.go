package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// if_acmpeq
// Branch if reference comparison succeeds
type IfACmpEq struct {
	base_instructions.BranchInstruction
}

func (ifACmpEq *IfACmpEq) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	referenceValue2 := operandStack.PopReferenceValue()
	referenceValue1 := operandStack.PopReferenceValue()

	if referenceValue1 == referenceValue2 {
		base_instructions.JumpToBranch(frame, ifACmpEq.Offset)
	}
}
