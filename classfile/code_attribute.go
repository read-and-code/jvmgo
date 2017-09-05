package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	constantPool              ConstantPool
	maxStackDepth             uint16
	maxNumberOfLocalVariables uint16
	code                      []byte
	exceptionTable            []*ExceptionTableEntry
	attributes                []AttributeInfo
}

func (codeAttribute *CodeAttribute) Read(classReader *ClassReader) {
	codeAttribute.maxStackDepth = classReader.ReadUint16()
	codeAttribute.maxNumberOfLocalVariables = classReader.ReadUint16()
	codeLength := classReader.ReadUint32()
	codeAttribute.code = classReader.ReadBytes(codeLength)
	codeAttribute.exceptionTable = readExceptionTable(classReader)
	codeAttribute.attributes = readAttributes(classReader, codeAttribute.constantPool)
}

func (codeAttribute *CodeAttribute) GetMaxStackDepth() uint {
	return uint(codeAttribute.maxStackDepth)
}

func (codeAttribute *CodeAttribute) GetMaxNumberOfLocalVariables() uint {
	return uint(codeAttribute.maxNumberOfLocalVariables)
}

func (codeAttribute *CodeAttribute) GetCode() []byte {
	return codeAttribute.code
}

func (codeAttribute *CodeAttribute) GetExceptionTable() []*ExceptionTableEntry {
	return codeAttribute.exceptionTable
}
