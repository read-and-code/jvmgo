package reference_instructions

import (
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// new
// Create new object.
// The unsigned indexbyte1 and indexbyte2 are used to construct an index into the run-time constant pool of the current class,
// where the value of the index is (indexbyte1 << 8) | indexbyte2.
// The run-time constant pool item at the index must be a symbolic reference to a class or interface type.
// The named class or interface type is resolved (§5.4.3.1) and should result in a class type.
// Memory for a new instance of that class is allocated from the garbage-collected heap,
// and the instance variables of the new object are initialized to their default initial values.
// The objectref, a reference to the instance, is pushed onto the operand stack.
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
