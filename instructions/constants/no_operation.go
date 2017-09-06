package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// nop
type NoOperation struct {
	base.NoOperandsInstruction
}

func (noOperation *NoOperation) Execute(frame *runtime_data_area.Frame) {
}
