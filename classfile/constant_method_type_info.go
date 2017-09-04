package classfile

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (constantMethodTypeInfo *ConstantMethodTypeInfo) Read(classReader *ClassReader) {
	constantMethodTypeInfo.descriptorIndex = classReader.ReadUint16()
}
