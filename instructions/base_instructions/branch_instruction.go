package base_instructions

type BranchInstruction struct {
	Offset int
}

func (branchInstruction *BranchInstruction) FetchOperands(bytecodeReader *BytecodeReader) {
	branchInstruction.Offset = int(bytecodeReader.ReadInt16())
}
