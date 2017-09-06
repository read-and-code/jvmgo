package classfile

import "fmt"

type ConstantPool []ConstantInfo

func readConstantPool(classReader *ClassReader) ConstantPool {
	constantPoolCount := int(classReader.ReadUint16())
	constantPool := make([]ConstantInfo, constantPoolCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1
	for i := 1; i < constantPoolCount; i++ {
		constantPool[i] = readConstantInfo(classReader, constantPool)

		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return constantPool
}

func (constantPool ConstantPool) GetConstantInfo(index uint16) ConstantInfo {
	constantInfo := constantPool[index]

	if constantInfo != nil {
		return constantInfo
	}

	panic(fmt.Errorf("Invalid constant pool index: %v", index))
}

func (constantPool ConstantPool) GetClassName(index uint16) string {
	constantClassReferenceInfo := constantPool.GetConstantInfo(index).(*ConstantClassReferenceInfo)

	return constantPool.GetUtf8String(constantClassReferenceInfo.nameIndex)
}

func (constantPool ConstantPool) GetNameAndTypeDescriptor(index uint16) (string, string) {
	constantNameAndTypeDescriptorInfo := constantPool.GetConstantInfo(index).(*ConstantNameAndTypeDescriptorInfo)
	name := constantPool.GetUtf8String(constantNameAndTypeDescriptorInfo.nameIndex)
	descriptor := constantPool.GetUtf8String(constantNameAndTypeDescriptorInfo.descriptorIndex)

	return name, descriptor
}

func (constantPool ConstantPool) GetUtf8String(index uint16) string {
	constantUtf8StringInfo := constantPool.GetConstantInfo(index).(*ConstantUtf8StringInfo)

	return constantUtf8StringInfo.value
}
