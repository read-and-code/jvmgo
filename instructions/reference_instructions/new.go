package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// new
// Create new object
type New struct {
	base_instructions.Index16Instruction
}

func (new *New) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classReference := constantPool.GetConstant(new.Index).(*heap.ClassReference)
	class := classReference.GetResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	frame.GetOperandStack().PushReferenceValue(class.NewObject())
}
