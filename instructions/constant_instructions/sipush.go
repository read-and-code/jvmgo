package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// sipush
// Push short. The immediate unsigned byte1 and byte2 values are assembled into an
// intermediate short where the value of the short is (byte1 << 8) | byte2.
// The intermediate value is then sign-extended to an int value.
// That value is pushed onto the operand stack.
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
