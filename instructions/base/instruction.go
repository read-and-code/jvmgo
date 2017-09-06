package base

import "github.com/Frederick-S/jvmgo/runtime_data_area"

type Instruction interface {
	FetchOperands(bytecodeReader *BytecodeReader)
	Execute(frame *runtime_data_area.Frame)
}
