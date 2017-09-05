package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (exceptionsAttribute *ExceptionsAttribute) Read(classReader *ClassReader) {
	exceptionsAttribute.exceptionIndexTable = classReader.ReadUint16Table()
}

func (exceptionsAttribute *ExceptionsAttribute) GetExceptionIndexTable() []uint16 {
	return exceptionsAttribute.exceptionIndexTable
}
