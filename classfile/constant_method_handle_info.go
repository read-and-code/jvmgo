package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	methodHandleKind           uint8
	methodHandleReferenceIndex uint16
}

func (constantMethodHandleInfo *ConstantMethodHandleInfo) Read(classReader *ClassReader) {
	constantMethodHandleInfo.methodHandleKind = classReader.ReadUint8()
	constantMethodHandleInfo.methodHandleReferenceIndex = classReader.ReadUint16()
}
