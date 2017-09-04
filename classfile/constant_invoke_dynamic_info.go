package classfile

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttributeIndex uint16
	nameAndTypeIndex              uint16
}

func (constantInvokeDynamicInfo *ConstantInvokeDynamicInfo) Read(classReader *ClassReader) {
	constantInvokeDynamicInfo.bootstrapMethodAttributeIndex = classReader.ReadUint16()
	constantInvokeDynamicInfo.nameAndTypeIndex = classReader.ReadUint16()
}
