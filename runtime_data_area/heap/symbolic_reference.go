package heap

type SymbolicReference struct {
	constantPool *ConstantPool
	className    string
	class        *Class
}

func (symbolicReference *SymbolicReference) GetResolvedClass() *Class {
	if symbolicReference.class == nil {
		symbolicReference.ResolveClassReference()
	}

	return symbolicReference.class
}

func (symbolicReference *SymbolicReference) ResolveClassReference() {
	class := symbolicReference.constantPool.class
	referencedClass := class.classLoader.LoadClass(symbolicReference.className)

	if !referencedClass.IsAccessibleTo(class) {
		panic("java.lang.IllegalAccessError")
	}

	symbolicReference.class = referencedClass
}
