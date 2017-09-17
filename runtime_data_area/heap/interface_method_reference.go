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
	class := interfaceMethodReference.constantPool.class
	resolvedClass := interfaceMethodReference.GetResolvedClass()

	if !resolvedClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(resolvedClass, interfaceMethodReference.name, interfaceMethodReference.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.IsAccessibleTo(class) {
		panic("java.lang.IllegalAccessError")
	}

	interfaceMethodReference.method = method
}

func lookupInterfaceMethod(interfaceMember *Class, name, descriptor string) *Method {
	for _, method := range interfaceMember.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInInterfaces(interfaceMember.interfaces, name, descriptor)
}
