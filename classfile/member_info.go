package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}

method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(classReader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := classReader.ReadUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members {
		members[i] = readMember(classReader, constantPool)
	}

	return members
}

func readMember(classReader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     classReader.ReadUint16(),
		nameIndex:       classReader.ReadUint16(),
		descriptorIndex: classReader.ReadUint16(),
		attributes:      readAttributes(classReader, constantPool),
	}
}

func (memberInfo *MemberInfo) GetAccessFlags() uint16 {
	return memberInfo.accessFlags
}

func (memberInfo *MemberInfo) GetName() string {
	return memberInfo.constantPool.GetUtf8String(memberInfo.nameIndex)
}

func (memberInfo *MemberInfo) GetDescriptor() string {
	return memberInfo.constantPool.GetUtf8String(memberInfo.descriptorIndex)
}
