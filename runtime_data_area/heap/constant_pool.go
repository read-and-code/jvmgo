package heap

import "github.com/Frederick-S/jvmgo/classfile"
import "fmt"

type ConstantPool struct {
	class     *Class
	constants []Constant
}

func newConstantPool(class *Class, classFileConstantPool classfile.ConstantPool) *ConstantPool {
	constantPoolSize := len(classFileConstantPool)
	constants := make([]Constant, constantPoolSize)
	constantPool := &ConstantPool{
		class,
		constants,
	}

	for i := 1; i < constantPoolSize; i++ {
		constantInfo := classFileConstantPool[i]

		switch constantInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			constants[i] = constantInfo.(*classfile.ConstantIntegerInfo).GetValue()
		case *classfile.ConstantFloatInfo:
			constants[i] = constantInfo.(*classfile.ConstantFloatInfo).GetValue()
		case *classfile.ConstantLongInfo:
			constants[i] = constantInfo.(*classfile.ConstantLongInfo).GetValue()
		case *classfile.ConstantDoubleInfo:
			constants[i] = constantInfo.(*classfile.ConstantDoubleInfo).GetValue()
		case *classfile.ConstantStringReferenceInfo:
			constants[i] = constantInfo.(*classfile.ConstantStringReferenceInfo).GetString()
		case *classfile.ConstantClassInfo:
			//
		case *classfile.ConstantFieldReferenceInfo:
			//
		case *classfile.ConstantMethodReferenceInfo:
			//
		case *classfile.ConstantInterfaceMethodReferenceInfo:
			//
		default:
			// pass
		}
	}

	return constantPool
}

func (constantPool *ConstantPool) GetConstant(index uint) Constant {
	constant := constantPool.constants[index]

	if constant != nil {
		return constant
	}

	panic(fmt.Sprintf("No constant at index %d", index))
}
