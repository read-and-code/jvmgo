package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// lookupswitch
// Access jump table by key match and jump.
type LookupSwitch struct {
	defaultOffset             int32
	numberOfMatchOffsetPairss int32
	matchOffsetPairss         []int32
}

func (lookupSwitch *LookupSwitch) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	bytecodeReader.SkipPadding()

	lookupSwitch.defaultOffset = bytecodeReader.ReadInt32()
	lookupSwitch.numberOfMatchOffsetPairss = bytecodeReader.ReadInt32()
	lookupSwitch.matchOffsetPairss = bytecodeReader.ReadInt32Table(lookupSwitch.numberOfMatchOffsetPairss * 2)
}

func (lookupSwitch *LookupSwitch) Execute(frame *runtime_data_area.Frame) {
	key := frame.GetOperandStack().PopIntegerValue()

	for i := int32(0); i < lookupSwitch.numberOfMatchOffsetPairss*2; i += 2 {
		if lookupSwitch.matchOffsetPairss[i] == key {
			offset := lookupSwitch.matchOffsetPairss[i+1]

			base_instructions.JumpToBranch(frame, int(offset))

			return
		}
	}

	base_instructions.JumpToBranch(frame, int(lookupSwitch.defaultOffset))
}
