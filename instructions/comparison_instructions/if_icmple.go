package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// if_icmple
// Branch if int comparison succeeds
type IfICmpLe struct {
	base_instructions.BranchInstruction
}

func (ifICmpLe *IfICmpLe) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue2 := operandStack.PopIntegerValue()
	integerValue1 := operandStack.PopIntegerValue()

	if integerValue1 <= integerValue2 {
		base_instructions.JumpToBranch(frame, ifICmpLe.Offset)
	}
}
