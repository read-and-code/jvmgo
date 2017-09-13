package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

type GoTo struct {
	base_instructions.BranchInstruction
}

func (goTo *GoTo) Execute(frame *runtime_data_area.Frame) {
	base_instructions.JumpToBranch(frame, goTo.Offset)
}
