package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// aconst_null
// Push the null object reference onto the operand stack.
type AConstNull struct {
	base_instructions.NoOperandsInstruction
}

func (aConstNull *AConstNull) Execute(frame *runtime_data_area.Frame) {
	frame.GetOperandStack().PushReferenceValue(nil)
}
