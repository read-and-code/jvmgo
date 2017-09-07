package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// bipush
type PushByteValue struct {
	value int8
}

func (pushByteValue *PushByteValue) FetchOperands(bytecodeReader *base.BytecodeReader) {
	pushByteValue.value = bytecodeReader.ReadInt8()
}

func (pushByteValue *PushByteValue) Execute(frame *runtime_data_area.Frame) {
	integer := int32(pushByteValue.value)

	frame.GetOperandStack().PushIntegerValue(integer)
}
