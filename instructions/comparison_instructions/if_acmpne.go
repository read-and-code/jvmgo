package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// if_acmpne
// Branch if reference comparison succeeds
type IfACmpNe struct {
	base_instructions.BranchInstruction
}

func (ifACmpNe *IfACmpNe) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	referenceValue2 := operandStack.PopReferenceValue()
	referenceValue1 := operandStack.PopReferenceValue()

	if referenceValue1 != referenceValue2 {
		base_instructions.JumpToBranch(frame, ifACmpNe.Offset)
	}
}
