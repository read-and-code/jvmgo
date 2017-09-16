package heap

import "github.com/Frederick-S/jvmgo/classfile"

type Field struct {
	ClassMember
	constantValueIndex uint
	variableIndex      uint
}

func newFields(class *Class, memberInfos []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(memberInfos))

	for i, memberInfo := range memberInfos {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(memberInfo)
		fields[i].copyAttributes(memberInfo)
	}

	return fields
}

func (field *Field) copyAttributes(memberInfo *classfile.MemberInfo) {
	constantValueAttribute := memberInfo.GetConstantValueAttribute()

	if constantValueAttribute != nil {
		field.constantValueIndex = uint(constantValueAttribute.GetConstantValueIndex())
	}
}

func (field *Field) IsVolatile() bool {
	return field.accessFlags&ACC_VOLATILE != 0
}

func (field *Field) IsTransient() bool {
	return field.accessFlags&ACC_TRANSIENT != 0
}

func (field *Field) IsEnum() bool {
	return field.accessFlags&ACC_ENUM != 0
}

func (field *Field) GetConstantValueIndex() uint {
	return field.constantValueIndex
}

func (field *Field) GetVariableIndex() uint {
	return field.variableIndex
}

func (field *Field) IsLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}
