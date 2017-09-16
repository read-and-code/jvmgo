package comparison_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// if_icmpge
// Branch if int comparison succeeds
type IfICmpGe struct {
	base_instructions.BranchInstruction
}

func (ifICmpGe *IfICmpGe) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	integerValue2 := operandStack.PopIntegerValue()
	integerValue1 := operandStack.PopIntegerValue()

	if integerValue1 >= integerValue2 {
		base_instructions.JumpToBranch(frame, ifICmpGe.Offset)
	}
}
