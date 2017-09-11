package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	constantPool ConstantPool
	nameIndex    uint16
}

func (constantClassInfo *ConstantClassInfo) Read(classReader *ClassReader) {
	constantClassInfo.nameIndex = classReader.ReadUint16()
}

func (constantClassInfo *ConstantClassInfo) GetName() string {
	return constantClassInfo.constantPool.GetUtf8String(constantClassInfo.nameIndex)
}
