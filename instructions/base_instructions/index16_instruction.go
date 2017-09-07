package base_instructions

type Index16Instruction struct {
	Index uint
}

func (index16Instruction *Index16Instruction) FetchOperands(bytecodeReader *BytecodeReader) {
	index16Instruction.Index = uint(bytecodeReader.ReadUint16())
}
