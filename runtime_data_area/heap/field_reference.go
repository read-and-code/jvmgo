package heap

import "github.com/Frederick-S/jvmgo/classfile"

type FieldReference struct {
	MemberReference
	field *Field
}

func newFieldReference(constantPool *ConstantPool, constantFieldReferenceInfo *classfile.ConstantFieldReferenceInfo) *FieldReference {
	fieldReference := &FieldReference{}
	fieldReference.constantPool = constantPool
	fieldReference.copyConstantMemberReferenceInfo(&constantFieldReferenceInfo.ConstantMemberReferenceInfo)

	return fieldReference
}

func (fieldReference *FieldReference) GetResolvedField() *Field {
	if fieldReference.field == nil {
		fieldReference.ResolveFieldReference()
	}

	return fieldReference.field
}

func (fieldReference *FieldReference) ResolveFieldReference() {
	class := fieldReference.constantPool.class
	resolvedClass := fieldReference.GetResolvedClass()
	field := lookupField(resolvedClass, fieldReference.name, fieldReference.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.IsAccessibleTo(class) {
		panic("java.lang.IllegalAccessError")
	}

	fieldReference.field = field
}

func lookupField(class *Class, name, descriptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, interfaceMember := range class.interfaces {
		field := lookupField(interfaceMember, name, descriptor)

		if field != nil {
			return field
		}
	}

	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}

	return nil
}
