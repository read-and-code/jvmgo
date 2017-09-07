package base_instructions

type BranchInstruction struct {
	offset int
}

func (branchInstruction *BranchInstruction) FetchOperands(bytecodeReader *BytecodeReader) {
	branchInstruction.offset = int(bytecodeReader.ReadInt16())
}
