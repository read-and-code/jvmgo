package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ldc_w
// Push item from run-time constant pool (wide index)
type LdcW struct {
	base_instructions.Index16Instruction
}

func (ldcW *LdcW) Execute(frame *runtime_data_area.Frame) {
	loadConstantFromConstantPoolAndPushToOperandStack(frame, ldcW.Index)
}
