package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iinc
// Increment local variable by constant
type IInc struct {
	Index    uint
	Constant int32
}

func (iInc *IInc) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	iInc.Index = uint(bytecodeReader.ReadUint8())
	iInc.Constant = int32(bytecodeReader.ReadInt8())
}

func (iInc *IInc) Execute(frame *runtime_data_area.Frame) {
	localVariables := frame.GetLocalVariables()
	value := localVariables.GetIntegerValue(iInc.Index)
	value += iInc.Constant

	localVariables.SetIntegerValue(iInc.Index, value)
}
