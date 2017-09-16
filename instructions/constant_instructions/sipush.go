package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// sipush
// Push short
type SIPush struct {
	value int16
}

func (sIPush *SIPush) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	sIPush.value = bytecodeReader.ReadInt16()
}

func (sIPush *SIPush) Execute(frame *runtime_data_area.Frame) {
	integer := int32(sIPush.value)

	frame.GetOperandStack().PushIntegerValue(integer)
}
