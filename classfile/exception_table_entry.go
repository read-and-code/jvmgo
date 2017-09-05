package classfile

type ExceptionTableEntry struct {
	startPC        uint16
	endPC          uint16
	handlerPC      uint16
	catchTypeIndex uint16
}

func readExceptionTable(classReader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := classReader.ReadUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPC:        classReader.ReadUint16(),
			endPC:          classReader.ReadUint16(),
			handlerPC:      classReader.ReadUint16(),
			catchTypeIndex: classReader.ReadUint16(),
		}
	}

	return exceptionTable
}

func (exceptionTableEntry *ExceptionTableEntry) GetStartPC() uint16 {
	return exceptionTableEntry.startPC
}

func (exceptionTableEntry *ExceptionTableEntry) GetEndPC() uint16 {
	return exceptionTableEntry.endPC
}

func (exceptionTableEntry *ExceptionTableEntry) GetHandlerPC() uint16 {
	return exceptionTableEntry.handlerPC
}

func (exceptionTableEntry *ExceptionTableEntry) GetCatchTypeIndex() uint16 {
	return exceptionTableEntry.catchTypeIndex
}
