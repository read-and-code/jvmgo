package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	Read(classReader *ClassReader)
}

func readAttributes(classReader *ClassReader, constantPool ConstantPool) []AttributeInfo {
	attributesCount := classReader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i := range attributes {
		attributes[i] = readAttribute(classReader, constantPool)
	}

	return attributes
}

func readAttribute(classReader *ClassReader, constantPool ConstantPool) AttributeInfo {
	attributeNameIndex := classReader.ReadUint16()
	attributeName := constantPool.GetUtf8String(attributeNameIndex)
	attributeLength := classReader.ReadUint32()
	attributeInfo := newAttributeInfo(attributeName, attributeLength, constantPool)
	attributeInfo.Read(classReader)

	return attributeInfo
}

func newAttributeInfo(attributeName string, attributeLength uint32, constantPool ConstantPool) AttributeInfo {
	switch attributeName {
	case "Code":
		return &CodeAttribute{constantPool: constantPool}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{constantPool: constantPool}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attributeName, attributeLength, nil}
	}
}
