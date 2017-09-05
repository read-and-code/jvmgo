package classfile

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPC         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (localVariableTableAttribute *LocalVariableTableAttribute) Read(classReader *ClassReader) {
	localVariableTableLength := classReader.ReadUint16()
	localVariableTable := make([]*LocalVariableTableEntry, localVariableTableLength)

	for i := range localVariableTable {
		localVariableTable[i] = &LocalVariableTableEntry{
			startPC:         classReader.ReadUint16(),
			length:          classReader.ReadUint16(),
			nameIndex:       classReader.ReadUint16(),
			descriptorIndex: classReader.ReadUint16(),
			index:           classReader.ReadUint16(),
		}
	}

	localVariableTableAttribute.localVariableTable = localVariableTable
}
