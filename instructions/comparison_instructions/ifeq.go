package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ifeq
// The value must be of type int.
// It is popped from the operand stack and compared against zero.
// All comparisons are signed. The results of the comparisons are as follows:
// ifeq succeeds if and only if value = 0
// If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset,
// where the offset is calculated to be (branchbyte1 << 8) | branchbyte2.
// Execution then proceeds at that offset from the address of the opcode of this if<cond> instruction.
// The target address must be that of an opcode of an instruction within the method that contains this if<cond> instruction.
// Otherwise, execution proceeds at the address of the instruction following this if<cond> instruction.
type IfEq struct {
	base_instructions.BranchInstruction
}

func (ifEq *IfEq) Execute(frame *runtime_data_area.Frame) {
	integerValue := frame.GetOperandStack().PopIntegerValue()

	if integerValue == 0 {
		base_instructions.JumpToBranch(frame, ifEq.Offset)
	}
}
