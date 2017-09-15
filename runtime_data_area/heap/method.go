package heap

import "github.com/Frederick-S/jvmgo/classfile"

type Method struct {
	ClassMember
	maxStackSize              uint
	maxNumberOfLocalVariables uint
	code                      []byte
}

func newMethods(class *Class, memberInfos []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(memberInfos))

	for i, memberInfo := range memberInfos {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(memberInfo)
		methods[i].copyAttributes(memberInfo)
	}

	return methods
}

func (method *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	codeAttribute := memberInfo.GetCodeAttribute()

	if codeAttribute != nil {
		method.maxStackSize = codeAttribute.GetMaxStackSize()
		method.maxNumberOfLocalVariables = codeAttribute.GetMaxNumberOfLocalVariables()
		method.code = codeAttribute.GetCode()
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
