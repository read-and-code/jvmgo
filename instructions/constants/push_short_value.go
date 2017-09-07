package constants

import (
	"github.com/Frederick-S/jvmgo/instructions/base"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// sipush
type PushShortValue struct {
	value int16
}

func (pushShortValue *PushShortValue) FetchOperands(bytecodeReader *base.BytecodeReader) {
	pushShortValue.value = bytecodeReader.ReadInt16()
}

func (pushShortValue *PushShortValue) Execute(frame *runtime_data_area.Frame) {
	integer := int32(pushShortValue.value)

	frame.GetOperandStack().PushIntegerValue(integer)
}
