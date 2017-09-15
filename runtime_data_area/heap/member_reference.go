package heap

import "github.com/Frederick-S/jvmgo/classfile"

type MemberReference struct {
	SymbolicReference
	name string
	descriptor string
}

func (memberReference *MemberReference) copyConstantMemberReferenceInfo(constantMemberReferenceInfo *classfile.ConstantMemberReferenceInfo) {
	memberReference.className = constantMemberReferenceInfo.GetClassName()
	memberReference.name, memberReference.descriptor = constantMemberReferenceInfo.GetNameAndTypeDescriptor()
}

func (memberReference *MemberReference) GetName() string {
	return memberReference.name
}

func (memberReference *MemberReference) GetDescriptor() string {
	return memberReference.descriptor
}