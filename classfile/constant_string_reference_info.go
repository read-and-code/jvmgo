package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringReferenceInfo struct {
	constantPool ConstantPool
	stringIndex  uint16
}

func (constantStringReferenceInfo *ConstantStringReferenceInfo) Read(classReader *ClassReader) {
	constantStringReferenceInfo.stringIndex = classReader.ReadUint16()
}

func (constantStringReferenceInfo *ConstantStringReferenceInfo) GetString() string {
	return constantStringReferenceInfo.constantPool.GetUtf8String(constantStringReferenceInfo.stringIndex)
}
