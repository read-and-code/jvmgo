package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassReferenceInfo struct {
	constantPool ConstantPool
	nameIndex    uint16
}

func (constantClassReferenceInfo *ConstantClassReferenceInfo) Read(classReader *ClassReader) {
	constantClassReferenceInfo.nameIndex = classReader.ReadUint16()
}

func (constantClassReferenceInfo *ConstantClassReferenceInfo) GetName() string {
	return constantClassReferenceInfo.constantPool.GetUtf8String(constantClassReferenceInfo.nameIndex)
}
