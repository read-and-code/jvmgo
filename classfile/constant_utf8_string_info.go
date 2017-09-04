package classfile

import (
	"fmt"
	"unicode/utf16"
)

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8StringInfo struct {
	value string
}

func (constantUtf8StringInfo *ConstantUtf8StringInfo) Read(classReader *ClassReader) {
	length := uint32(classReader.ReadUint16())
	bytes := classReader.ReadBytes(length)
	constantUtf8StringInfo.value = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	utfLength := len(bytes)
	chars := make([]uint16, utfLength)

	var char1, char2, char3 uint16
	count := 0
	charsCount := 0

	for count < utfLength {
		char1 = uint16(bytes[count])

		if char1 > 127 {
			break
		}

		count++
		chars[charsCount] = char1
		charsCount++
	}

	for count < utfLength {
		char1 = uint16(bytes[count])

		switch char1 >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chars[charsCount] = char1
			charsCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2

			if count > utfLength {
				panic("Malformed input: partial character at end")
			}

			char2 = uint16(bytes[count-1])

			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("Malformed input around byte %v", count))
			}

			chars[charsCount] = char1&0x1F<<6 | char2&0x3F
			charsCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3

			if count > utfLength {
				panic("Malformed input: partial character at end")
			}

			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])

			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("Malformed input around byte %v", (count - 1)))
			}

			chars[charsCount] = char1&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charsCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("Malformed input around byte %v", count))
		}
	}

	// The number of chars produced may be less than utflen
	chars = chars[0:charsCount]

	return string(utf16.Decode(chars))
}
