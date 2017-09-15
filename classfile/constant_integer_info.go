package classfile

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	value int32
}

func (constantIntegerInfo *ConstantIntegerInfo) Read(classReader *ClassReader) {
	bytes := classReader.ReadUint32()
	constantIntegerInfo.value = int32(bytes)
}

func (constantIntegerInfo *ConstantIntegerInfo) GetValue() int32 {
	return constantIntegerInfo.value
}
