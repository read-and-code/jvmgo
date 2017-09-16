package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ldc2_w
// Push long or double from run-time constant pool (wide index)
type Ldc2W struct {
	base_instructions.Index16Instruction
}

func (ldc2W *Ldc2W) Execute(frame *runtime_data_area.Frame) {
	operandStack := frame.GetOperandStack()
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	constant := constantPool.GetConstant(ldc2W.Index)

	switch constant.(type) {
	case int64:
		operandStack.PushLongValue(constant.(int64))
	case float64:
		operandStack.PushDoubleValue(constant.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
