package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	constantPool        ConstantPool
	sourceFileNameIndex uint16
}

func (sourceFileAttribute *SourceFileAttribute) Read(classReader *ClassReader) {
	sourceFileAttribute.sourceFileNameIndex = classReader.ReadUint16()
}

func (sourceFileAttribute *SourceFileAttribute) GetFileName() string {
	return sourceFileAttribute.constantPool.GetUtf8String(sourceFileAttribute.sourceFileNameIndex)
}
