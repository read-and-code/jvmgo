package classfile

type ConstantMemberReferenceInfo struct {
	constantPool     ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (constantMemberReferenceInfo *ConstantMemberReferenceInfo) Read(classReader *ClassReader) {
	constantMemberReferenceInfo.classIndex = classReader.ReadUint16()
	constantMemberReferenceInfo.nameAndTypeIndex = classReader.ReadUint16()
}

func (constantMemberReferenceInfo *ConstantMemberReferenceInfo) GetClassName() string {
	return constantMemberReferenceInfo.constantPool.GetClassName(constantMemberReferenceInfo.classIndex)
}

func (constantMemberReferenceInfo *ConstantMemberReferenceInfo) GetNameAndTypeDescriptor() (string, string) {
	return constantMemberReferenceInfo.constantPool.GetNameAndTypeDescriptor(constantMemberReferenceInfo.nameAndTypeIndex)
}
