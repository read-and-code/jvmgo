package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// u1
func (classReader *ClassReader) ReadUint8() uint8 {
	value := classReader.data[0]
	classReader.data = classReader.data[1:]

	return value
}

// u2
func (classReader *ClassReader) ReadUint16() uint16 {
	value := binary.BigEndian.Uint16(classReader.data)
	classReader.data = classReader.data[2:]

	return value
}

// u4
func (classReader *ClassReader) ReadUint32() uint32 {
	value := binary.BigEndian.Uint32(classReader.data)
	classReader.data = classReader.data[4:]

	return value
}

func (classReader *ClassReader) ReadUint64() uint64 {
	value := binary.BigEndian.Uint64(classReader.data)
	classReader.data = classReader.data[8:]

	return value
}

func (classReader *ClassReader) ReadUint16Table() []uint16 {
	numberOfItems := classReader.ReadUint16()
	uint16Table := make([]uint16, numberOfItems)

	for i := range uint16Table {
		uint16Table[i] = classReader.ReadUint16()
	}

	return uint16Table
}

func (classReader *ClassReader) ReadBytes(numberOfBytes uint32) []byte {
	bytes := classReader.data[:numberOfBytes]
	classReader.data = classReader.data[numberOfBytes:]

	return bytes
}
