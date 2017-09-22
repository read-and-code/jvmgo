package classfile

import "fmt"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magicNumber      uint32
	minorVersion     uint16
	majorVersion     uint16
	constantPool     ConstantPool
	accessFlags      uint16
	thisClassIndex   uint16
	superClassIndex  uint16
	interfaceIndices []uint16
	fields           []*MemberInfo
	methods          []*MemberInfo
	attributes       []AttributeInfo
}

func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		r := recover()

		if r != nil {
			var isOk bool

			err, isOk = r.(error)

			if !isOk {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	classReader := &ClassReader{classData}
	classFile = &ClassFile{}

	classFile.Read(classReader)

	return
}

func (classFile *ClassFile) Read(classReader *ClassReader) {
	classFile.ReadAndCheckMagicNumber(classReader)
	classFile.ReadAndCheckVersion(classReader)
	classFile.constantPool = readConstantPool(classReader)
	classFile.accessFlags = classReader.ReadUint16()
	classFile.thisClassIndex = classReader.ReadUint16()
	classFile.superClassIndex = classReader.ReadUint16()
	classFile.interfaceIndices = classReader.ReadUint16Table()
	classFile.fields = readMembers(classReader, classFile.constantPool)
	classFile.methods = readMembers(classReader, classFile.constantPool)
	classFile.attributes = readAttributes(classReader, classFile.constantPool)
}

func (classFile *ClassFile) ReadAndCheckMagicNumber(classReader *ClassReader) {
	magicNumber := classReader.ReadUint32()

	if magicNumber != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: invalid magic number!")
	}
}

func (classFile *ClassFile) ReadAndCheckVersion(classReader *ClassReader) {
	classFile.minorVersion = classReader.ReadUint16()
	classFile.majorVersion = classReader.ReadUint16()

	switch classFile.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if classFile.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (classFile *ClassFile) GetMinorVersion() uint16 {
	return classFile.minorVersion
}

func (classFile *ClassFile) GetMajorVersion() uint16 {
	return classFile.majorVersion
}

func (classFile *ClassFile) GetConstantPool() ConstantPool {
	return classFile.constantPool
}

func (classFile *ClassFile) GetAccessFlags() uint16 {
	return classFile.accessFlags
}

func (classFile *ClassFile) GetFields() []*MemberInfo {
	return classFile.fields
}

func (classFile *ClassFile) GetMethods() []*MemberInfo {
	return classFile.methods
}

func (classFile *ClassFile) GetClassName() string {
	return classFile.constantPool.GetClassName(classFile.thisClassIndex)
}

func (classFile *ClassFile) GetSuperClassName() string {
	if classFile.superClassIndex > 0 {
		return classFile.constantPool.GetClassName(classFile.superClassIndex)
	}

	return ""
}

func (classFile *ClassFile) GetInterfaceNames() []string {
	interfaceNames := make([]string, len(classFile.interfaceIndices))

	for i, interfaceIndex := range classFile.interfaceIndices {
		interfaceNames[i] = classFile.constantPool.GetClassName(interfaceIndex)
	}

	return interfaceNames
}

func (classFile *ClassFile) GetSourceFileAttribute() *SourceFileAttribute {
	for _, attributeInfo := range classFile.attributes {
		switch attributeInfo.(type) {
		case *SourceFileAttribute:
			return attributeInfo.(*SourceFileAttribute)
		}
	}

	return nil
}
