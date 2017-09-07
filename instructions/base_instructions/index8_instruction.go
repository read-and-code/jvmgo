package base_instructions

type Index8Instruction struct {
	Index uint
}

func (index8Instruction *Index8Instruction) FetchOperands(bytecodeReader *BytecodeReader) {
	index8Instruction.Index = uint(bytecodeReader.ReadUint8())
}
