package heap

import "github.com/Frederick-S/jvmgo/classfile"

type Method struct {
	ClassMember
	maxStackSize              uint
	maxNumberOfLocalVariables uint
	code                      []byte
	exceptionTable            ExceptionTable
	lineNumberTable           *classfile.LineNumberTableAttribute
	argumentsCount            uint
}

func newMethods(class *Class, memberInfos []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(memberInfos))

	for i, memberInfo := range memberInfos {
		methods[i] = newMethod(class, memberInfo)
	}

	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	methodDescriptor := parseMethodDescriptor(method.descriptor)
	method.calculateArgumentsCount(methodDescriptor.parameterTypes)

	if method.IsNative() {
		method.injectCodeAttribute(methodDescriptor.returnType)
	}

	return method
}

func (method *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	codeAttribute := memberInfo.GetCodeAttribute()

	if codeAttribute != nil {
		method.maxStackSize = codeAttribute.GetMaxStackSize()
		method.maxNumberOfLocalVariables = codeAttribute.GetMaxNumberOfLocalVariables()
		method.code = codeAttribute.GetCode()
		method.exceptionTable = newExceptionTable(codeAttribute.GetExceptionTable(), method.class.constantPool)
		method.lineNumberTable = codeAttribute.GetLineNumberTableAttribute()
	}
}

func (method *Method) calculateArgumentsCount(parameterTypes []string) {
	for _, parameterType := range parameterTypes {
		if parameterType == "J" || parameterType == "D" {
			method.argumentsCount += 2
		} else {
			method.argumentsCount += 1
		}
	}

	if !method.IsStatic() {
		// `this` reference
		method.argumentsCount++
	}
}

func (method *Method) injectCodeAttribute(returnType string) {
	method.maxStackSize = 4
	method.maxNumberOfLocalVariables = method.argumentsCount

	switch returnType[0] {
	case 'V':
		// return
		method.code = []byte{0xfe, 0xb1}
	case 'L', '[':
		// areturn
		method.code = []byte{0xfe, 0xb0}
	case 'D':
		// dreturn
		method.code = []byte{0xfe, 0xaf}
	case 'F':
		// freturn
		method.code = []byte{0xfe, 0xae}
	case 'J':
		// lreturn
		method.code = []byte{0xfe, 0xad}
	default:
		// ireturn
		method.code = []byte{0xfe, 0xac}
	}
}

func (method *Method) IsSynchronized() bool {
	return method.accessFlags&ACC_SYNCHRONIZED != 0
}

func (method *Method) IsBridge() bool {
	return method.accessFlags&ACC_BRIDGE != 0
}

func (method *Method) IsVarargs() bool {
	return method.accessFlags&ACC_VARARGS != 0
}

func (method *Method) IsNative() bool {
	return method.accessFlags&ACC_NATIVE != 0
}

func (method *Method) IsAbstract() bool {
	return method.accessFlags&ACC_ABSTRACT != 0
}

func (method *Method) IsStrict() bool {
	return method.accessFlags&ACC_STRICT != 0
}

func (method *Method) GetMaxStackSize() uint {
	return method.maxStackSize
}

func (method *Method) GetMaxNumberOfLocalVariables() uint {
	return method.maxNumberOfLocalVariables
}

func (method *Method) GetCode() []byte {
	return method.code
}

func (method *Method) GetArgumentsCount() uint {
	return method.argumentsCount
}

func (method *Method) FindExceptionHandler(exceptionClass *Class, pc int) int {
	exceptionHandler := method.exceptionTable.findExceptionHandler(exceptionClass, pc)

	if exceptionHandler != nil {
		return exceptionHandler.handlerPC
	}

	return -1
}

func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}

	if method.lineNumberTable == nil {
		return -1
	}

	return method.lineNumberTable.GetLineNumber(pc)
}
