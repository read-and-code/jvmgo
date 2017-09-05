package classfile

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPC    uint16
	lineNumber uint16
}

func (lineNumberTableAttribute *LineNumberTableAttribute) Read(classReader *ClassReader) {
	lineNumberTableLength := classReader.ReadUint16()
	lineNumberTable := make([]*LineNumberTableEntry, lineNumberTableLength)

	for i := range lineNumberTable {
		lineNumberTable[i] = &LineNumberTableEntry{
			startPC:    classReader.ReadUint16(),
			lineNumber: classReader.ReadUint16(),
		}
	}

	lineNumberTableAttribute.lineNumberTable = lineNumberTable
}

func (lineNumberTableAttribute *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lineNumberTableAttribute.lineNumberTable) - 1; i >= 0; i-- {
		lineNumberTableEntry := lineNumberTableAttribute.lineNumberTable[i]

		if pc >= int(lineNumberTableEntry.startPC) {
			return int(lineNumberTableEntry.lineNumber)
		}
	}

	return -1
}
