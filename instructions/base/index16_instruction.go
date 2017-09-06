package base

type Index16Instruction struct {
	index uint
}

func (index16Instruction *Index16Instruction) FetchOperands(bytecodeReader *BytecodeReader) {
	index16Instruction.index = uint(bytecodeReader.ReadUint16())
}
