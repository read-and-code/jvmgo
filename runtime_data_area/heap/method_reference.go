package heap

import "github.com/Frederick-S/jvmgo/classfile"

type MethodReference struct {
	MemberReference
	method *Method
}

func newMethodReference(constantPool *ConstantPool, constantMethodReferenceInfo *classfile.ConstantMethodReferenceInfo) *MethodReference {
	methodReference := &MethodReference{}
	methodReference.constantPool = constantPool
	methodReference.copyConstantMemberReferenceInfo(&constantMethodReferenceInfo.ConstantMemberReferenceInfo)

	return methodReference
}

func (methodReference *MethodReference) GetResolvedMethod() *Method {
	if methodReference.method == nil {
		methodReference.ResolveMethodReference()
	}

	return methodReference.method
}

func (methodReference *MethodReference) ResolveMethodReference() {
}
