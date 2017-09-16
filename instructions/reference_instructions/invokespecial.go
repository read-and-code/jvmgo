package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// invokespecial
// Invoke instance method; special handling for superclass, private, and instance initialization method invocations
type InvokeSpecial struct {
	base_instructions.Index16Instruction
}

func (invokeSpecial *InvokeSpecial) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PopReferenceValue()
}
