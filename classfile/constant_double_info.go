package classfile

import "math"

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	value float64
}

func (constantDoubleInfo *ConstantDoubleInfo) Read(classReader *ClassReader) {
	bytes := classReader.ReadUint64()
	constantDoubleInfo.value = math.Float64frombits(bytes)
}

func (constantDoubleInfo *ConstantDoubleInfo) GetValue() float64 {
	return constantDoubleInfo.value
}
