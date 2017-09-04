package classfile

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	value int64
}

func (constantLongInfo *ConstantLongInfo) Read(classReader *ClassReader) {
	bytes := classReader.ReadUint64()
	constantLongInfo.value = int64(bytes)
}
