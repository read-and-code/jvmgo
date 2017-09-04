package classfile

// Constant pool tags
// https://en.wikipedia.org/wiki/Java_class_file#The_constant_pool
const (
	constantTypeClassReference           = 7
	constantTypeFieldReference           = 9
	constantTypeMethodReference          = 10
	constantTypeInterfaceMethodReference = 11
	constantTypeStringReference          = 8
	constantTypeInteger                  = 3
	constantTypeFloat                    = 4
	constantTypeLong                     = 5
	constantTypeDouble                   = 6
	constantTypeNameAndTypeDescriptor    = 12
	constantTypeUtf8String               = 1
	constantTypeMethodHandle             = 15
	constantTypeMethodType               = 16
	constantTypeInvokeDynamic            = 18
)

type ConstantInfo interface {
	Read(classReader *ClassReader)
}

func readConstantInfo(classReader *ClassReader, constantPool ConstantPool) ConstantInfo {
	constantInfoType := classReader.ReadUint8()
	constantInfo := newConstantInfo(constantInfoType, constantPool)
	constantInfo.Read(classReader)

	return constantInfo
}

func newConstantInfo(constantInfoType uint8, constantPool ConstantPool) ConstantInfo {
	switch constantInfoType {
	case constantTypeInteger:
		return &ConstantIntegerInfo{}
	case constantTypeFloat:
		return &ConstantFloatInfo{}
	case constantTypeLong:
		return &ConstantLongInfo{}
	case constantTypeDouble:
		return &ConstantDoubleInfo{}
	case constantTypeUtf8String:
		return &ConstantUtf8StringInfo{}
	case constantTypeStringReference:
		return &ConstantStringReferenceInfo{constantPool: constantPool}
	case constantTypeClassReference:
		return &ConstantClassReferenceInfo{constantPool: constantPool}
	case constantTypeFieldReference:
		return &ConstantFieldReferenceInfo{ConstantMemberReferenceInfo{constantPool: constantPool}}
	case constantTypeMethodReference:
		return &ConstantMethodReferenceInfo{ConstantMemberReferenceInfo{constantPool: constantPool}}
	case constantTypeInterfaceMethodReference:
		return &ConstantInterfaceMethodReferenceInfo{ConstantMemberReferenceInfo{constantPool: constantPool}}
	case constantTypeNameAndTypeDescriptor:
		return &ConstantNameAndTypeDescriptorInfo{}
	case constantTypeMethodType:
		return &ConstantMethodTypeInfo{}
	case constantTypeMethodHandle:
		return &ConstantMethodHandleInfo{}
	case constantTypeInvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: unsupported constant pool tag!")
	}
}
