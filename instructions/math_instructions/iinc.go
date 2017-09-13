package math_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// iinc
// Increment local variable by constant.
// The index is an unsigned byte that must be an index into the local variable array of the current frame.
// The const is an immediate signed byte.
// The local variable at index must contain an int.
// The value const is first sign-extended to an int, and then the local variable at index is incremented by that amount.
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
