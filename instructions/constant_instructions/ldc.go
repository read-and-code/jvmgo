package constant_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

// ldc
// Push item from run-time constant pool
type Ldc struct {
	base_instructions.Index8Instruction
}

func (ldc *Ldc) Execute(frame *runtime_data_area.Frame) {
	loadConstantFromConstantPoolAndPushToOperandStack(frame, ldc.Index)
}

func loadConstantFromConstantPoolAndPushToOperandStack(frame *runtime_data_area.Frame, index uint) {
	operandStack := frame.GetOperandStack()
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	constant := constantPool.GetConstant(index)

	switch constant.(type) {
	case int32:
		operandStack.PushIntegerValue(constant.(int32))
	case float32:
		operandStack.PushFloatValue(constant.(float32))
	default:
		panic("TODO: ldc")
	}
}
