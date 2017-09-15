package heap

import "github.com/Frederick-S/jvmgo/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (classMember *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	classMember.accessFlags = memberInfo.GetAccessFlags()
	classMember.name = memberInfo.GetName()
	classMember.descriptor = memberInfo.GetDescriptor()
}

func (classMember *ClassMember) IsPublic() bool {
	return classMember.accessFlags&ACC_PUBLIC != 0
}

func (classMember *ClassMember) IsPrivate() bool {
	return classMember.accessFlags&ACC_PRIVATE != 0
}

func (classMember *ClassMember) IsProtected() bool {
	return classMember.accessFlags&ACC_PROTECTED != 0
}

func (classMember *ClassMember) IsStatic() bool {
	return classMember.accessFlags&ACC_STATIC != 0
}

func (classMember *ClassMember) IsFinal() bool {
	return classMember.accessFlags&ACC_FINAL != 0
}

func (classMember *ClassMember) IsSynthetic() bool {
	return classMember.accessFlags&ACC_SYNTHETIC != 0
}

func (classMember *ClassMember) GetName() string {
	return classMember.GetName()
}

func (classMember *ClassMember) GetDescriptor() string {
	return classMember.descriptor
}

func (classMember *ClassMember) GetClass() *Class {
	return classMember.class
}

func (classMember *ClassMember) IsAccessibleTo(otherClass *Class) bool {
	if classMember.IsPublic() {
		return true
	}

	class := classMember.class

	if classMember.IsProtected() {
		return otherClass == class || otherClass.IsSubClassOf(class) || class.GetPackageName() == otherClass.GetPackageName()
	}

	if !classMember.IsPrivate() {
		return class.GetPackageName() == otherClass.GetPackageName()
	}

	return class == otherClass
}
