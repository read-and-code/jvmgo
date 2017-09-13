package extended_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/instructions/load_instructions"
	"github.com/Frederick-S/jvmgo/instructions/math_instructions"
	"github.com/Frederick-S/jvmgo/instructions/store_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// wide
// Extend local variable index by additional bytes.
type Wide struct {
	modifiedInstruction base_instructions.Instruction
}

func (wide *Wide) FetchOperands(bytecodeReader *base_instructions.BytecodeReader) {
	operationCode := bytecodeReader.ReadUint8()

	switch operationCode {
	case 0x15:
		instruction := &load_instructions.ILoad{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x16:
		instruction := &load_instructions.LLoad{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x17:
		instruction := &load_instructions.FLoad{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x18:
		instruction := &load_instructions.DLoad{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x19:
		instruction := &load_instructions.ALoad{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x36:
		instruction := &store_instructions.IStore{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x37:
		instruction := &store_instructions.LStore{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x38:
		instruction := &store_instructions.FStore{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x39:
		instruction := &store_instructions.DStore{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x3a:
		instruction := &store_instructions.AStore{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		wide.modifiedInstruction = instruction
	case 0x84:
		instruction := &math_instructions.IInc{}
		instruction.Index = uint(bytecodeReader.ReadUint16())
		instruction.Constant = int32(bytecodeReader.ReadInt16())
		wide.modifiedInstruction = instruction
	case 0xa9:
		panic("Unsupported operation code: 0xa9!")
	}
}

func (wide *Wide) Execute(frame *runtime_data_area.Frame) {
	wide.modifiedInstruction.Execute(frame)
}
