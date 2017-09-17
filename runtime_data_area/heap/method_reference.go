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
	class := methodReference.constantPool.class
	resolvedClass := methodReference.GetResolvedClass()

	if resolvedClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(resolvedClass, methodReference.name, methodReference.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.IsAccessibleTo(class) {
		panic("java.lang.IllegalAccessError")
	}

	methodReference.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)

	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}

	return method
}
