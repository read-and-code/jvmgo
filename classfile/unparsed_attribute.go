package classfile

type UnparsedAttribute struct {
	name       string
	dataLength uint32
	data       []byte
}

func (unparsedAttribute *UnparsedAttribute) Read(classReader *ClassReader) {
	unparsedAttribute.data = classReader.ReadBytes(unparsedAttribute.dataLength)
}

func (unparsedAttribute *UnparsedAttribute) GetData() []byte {
	return unparsedAttribute.data
}
