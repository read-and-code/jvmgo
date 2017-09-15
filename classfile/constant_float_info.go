package classfile

import "math"

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	value float32
}

func (constantFloatInfo *ConstantFloatInfo) Read(classReader *ClassReader) {
	bytes := classReader.ReadUint32()
	constantFloatInfo.value = math.Float32frombits(bytes)
}

func (constantFloatInfo *ConstantFloatInfo) GetValue() float32 {
	return constantFloatInfo.value
}
