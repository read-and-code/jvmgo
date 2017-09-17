package heap

import (
	"strings"

	"github.com/Frederick-S/jvmgo/classfile"
)

type Class struct {
	accessFlags            uint16
	name                   string
	superClassName         string
	interfaceNames         []string
	constantPool           *ConstantPool
	fields                 []*Field
	methods                []*Method
	classLoader            *ClassLoader
	superClass             *Class
	interfaces             []*Class
	instanceVariablesCount uint
	staticVariablesCount   uint
	staticVariables        Variables
}

func newClass(classFile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classFile.GetAccessFlags()
	class.name = classFile.GetClassName()
	class.superClassName = classFile.GetSuperClassName()
	class.interfaceNames = classFile.GetInterfaceNames()
	class.constantPool = newConstantPool(class, classFile.GetConstantPool())
	class.fields = newFields(class, classFile.GetFields())
	class.methods = newMethods(class, classFile.GetMethods())

	return class
}

func (class *Class) IsPublic() bool {
	return class.accessFlags&ACC_PUBLIC != 0
}

func (class *Class) IsFinal() bool {
	return class.accessFlags&ACC_FINAL != 0
}

func (class *Class) IsSuper() bool {
	return class.accessFlags&ACC_SUPER != 0
}

func (class *Class) IsInterface() bool {
	return class.accessFlags&ACC_INTERFACE != 0
}

func (class *Class) IsAbstract() bool {
	return class.accessFlags&ACC_ABSTRACT != 0
}

func (class *Class) IsSynthetic() bool {
	return class.accessFlags&ACC_SYNTHETIC != 0
}

func (class *Class) IsAnnotation() bool {
	return class.accessFlags&ACC_ANNOTATION != 0
}

func (class *Class) IsEnum() bool {
	return class.accessFlags&ACC_ENUM != 0
}

func (class *Class) GetName() string {
	return class.name
}

func (class *Class) GetConstantPool() *ConstantPool {
	return class.constantPool
}

func (class *Class) IsAccessibleTo(otherClass *Class) bool {
	return class.IsPublic() || class.GetPackageName() == otherClass.GetPackageName()
}

func (class *Class) IsAssignableFrom(otherClass *Class) bool {
	if class == otherClass {
		return true
	}

	if !class.IsInterface() {
		return otherClass.IsSubClassOf(class)
	} else {
		return otherClass.IsImplementsFrom(class)
	}
}

func (class *Class) IsSubClassOf(otherClass *Class) bool {
	for currentClass := class; currentClass != nil; currentClass = currentClass.superClass {
		if currentClass == otherClass {
			return true
		}
	}

	return false
}

func (class *Class) IsImplementsFrom(otherInterface *Class) bool {
	for currentClass := class; currentClass != nil; currentClass = currentClass.superClass {
		for _, interfaceMember := range currentClass.interfaces {
			if interfaceMember == otherInterface || interfaceMember.IsSubInterfaceOf(otherInterface) {
				return true
			}
		}
	}

	return false
}

func (class *Class) IsSubInterfaceOf(otherInterface *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == otherInterface || superInterface.IsSubInterfaceOf(otherInterface) {
			return true
		}
	}

	return false
}

func (class *Class) GetPackageName() string {
	index := strings.LastIndex(class.name, "/")

	if index >= 0 {
		return class.name[:index]
	}

	return ""
}

func (class *Class) GetMainMethod() *Method {
	return class.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) GetStaticMethod(methodName, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() && method.name == methodName && method.descriptor == descriptor {
			return method
		}
	}

	return nil
}

func (class *Class) GetStaticVariables() Variables {
	return class.staticVariables
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}
