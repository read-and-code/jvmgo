package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// if_acmpne
// Branch if reference comparison succeeds.
// Both value1 and value2 must be of type reference.
// They are both popped from the operand stack and compared.
// The results of the comparison are as follows:
// if_acmpne succeeds if and only if value1 ≠ value2
// If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset,
// where the offset is calculated to be (branchbyte1 << 8) | branchbyte2.
// Execution then proceeds at that offset from the address of the opcode of this if_acmp<cond> instruction.
// The target address must be that of an opcode of an instruction within the method that contains this if_acmp<cond> instruction.
// Otherwise, if the comparison fails, execution proceeds at the address of the instruction following this if_acmp<cond> instruction.
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
