package base

type Index8Instruction struct {
	index uint
}

func (index8Instruction *Index8Instruction) FetchOperands(bytecodeReader *BytecodeReader) {
	index8Instruction.index = uint(bytecodeReader.ReadUint8())
}
