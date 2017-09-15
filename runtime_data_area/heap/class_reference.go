package heap

import "github.com/Frederick-S/jvmgo/classfile"

type ClassReference struct {
	SymbolicReference
}

func newClassReference(constantPool *ConstantPool, constantClassInfo *classfile.ConstantClassInfo) *ClassReference {
	classReference := &ClassReference{}
	classReference.constantPool = constantPool
	classReference.className = constantClassInfo.GetName()

	return classReference
}
