package heap

import "github.com/Frederick-S/jvmgo/classfile"

type InterfaceMethodReference struct {
	MemberReference
	method *Method
}

func newInterfaceMethodReference(constantPool *ConstantPool, constantInterfaceMethodReferenceInfo *classfile.ConstantInterfaceMethodReferenceInfo) *InterfaceMethodReference {
	interfaceMethodReference := &InterfaceMethodReference{}
	interfaceMethodReference.constantPool = constantPool
	interfaceMethodReference.copyConstantMemberReferenceInfo(&constantInterfaceMethodReferenceInfo.ConstantMemberReferenceInfo)

	return interfaceMethodReference
}

func (interfaceMethodReference *InterfaceMethodReference) GetResolvedInterfaceMethod() *Method {
	if interfaceMethodReference.method == nil {
		interfaceMethodReference.ResolveInterfaceMethodReference()
	}

	return interfaceMethodReference.method
}

func (interfaceMethodReference *InterfaceMethodReference) ResolveInterfaceMethodReference() {
}
