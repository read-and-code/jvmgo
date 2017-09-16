package control_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// tableswitch
// Access jump table by index and jump
type TableSwitch struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (tableSwitch *TableSwitch) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	bytecodeReader.SkipPadding()

	tableSwitch.defaultOffset = bytecodeReader.ReadInt32()
	tableSwitch.low = bytecodeReader.ReadInt32()
	tableSwitch.high = bytecodeReader.ReadInt32()
	tableSwitch.jumpOffsets = bytecodeReader.ReadInt32Table(tableSwitch.high - tableSwitch.low + 1)
}

func (tableSwitch *TableSwitch) Execute(frame *runtime_data_area.Frame) {
	index := frame.GetOperandStack().PopIntegerValue()

	var offset int

	if index >= tableSwitch.low && index <= tableSwitch.high {
		offset = int(tableSwitch.jumpOffsets[index-tableSwitch.low])
	} else {
		offset = int(tableSwitch.defaultOffset)
	}

	base_instructions.JumpToBranch(frame, offset)
}
