package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (constantValueAttribute *ConstantValueAttribute) Read(classReader *ClassReader) {
	constantValueAttribute.constantValueIndex = classReader.ReadUint16()
}

func (constantValueAttribute *ConstantValueAttribute) GetConstantValueIndex() uint16 {
	return constantValueAttribute.constantValueIndex
}
