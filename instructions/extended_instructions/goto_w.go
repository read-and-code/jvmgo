package extended_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// goto_w
// Branch always (wide index).
type GoToW struct {
	offset int
}

func (goToW *GoToW) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	goToW.offset = int(bytecodeReader.ReadInt32())
}

func (goToW *GoToW) Execute(frame *runtime_data_area.Frame) {
	base_instructions.JumpToBranch(frame, goToW.offset)
}
