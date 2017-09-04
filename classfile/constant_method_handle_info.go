package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (constantMethodHandleInfo *ConstantMethodHandleInfo) Read(classReader *ClassReader) {
	constantMethodHandleInfo.referenceKind = classReader.ReadUint8()
	constantMethodHandleInfo.referenceIndex = classReader.ReadUint16()
}
