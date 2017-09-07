package base_instructions

type BytecodeReader struct {
	code []byte
	pc   int
}

func (bytecodeReader *BytecodeReader) Reset(code []byte, pc int) {
	bytecodeReader.code = code
	bytecodeReader.pc = pc
}

func (bytecodeReader *BytecodeReader) GetPC() int {
	return bytecodeReader.pc
}

func (bytecodeReader *BytecodeReader) ReadInt8() int8 {
	return int8(bytecodeReader.ReadUint8())
}

func (bytecodeReader *BytecodeReader) ReadUint8() uint8 {
	integer := bytecodeReader.code[bytecodeReader.pc]
	bytecodeReader.pc++

	return integer
}

func (bytecodeReader *BytecodeReader) ReadInt16() int16 {
	return int16(bytecodeReader.ReadUint16())
}

func (bytecodeReader *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(bytecodeReader.ReadUint8())
	byte2 := uint16(bytecodeReader.ReadUint8())

	return (byte1 << 8) | byte2
}

func (bytecodeReader *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(bytecodeReader.ReadUint8())
	byte2 := int32(bytecodeReader.ReadUint8())
	byte3 := int32(bytecodeReader.ReadUint8())
	byte4 := int32(bytecodeReader.ReadUint8())

	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (bytecodeReader *BytecodeReader) ReadInt32Table(count int32) []int32 {
	int32Table := make([]int32, count)

	for i := range int32Table {
		int32Table[i] = bytecodeReader.ReadInt32()
	}

	return int32Table
}

func (bytecodeReader *BytecodeReader) SkipPadding() {
	for bytecodeReader.pc%4 != 0 {
		bytecodeReader.ReadUint8()
	}
}
