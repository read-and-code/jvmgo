package reference_instructions

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

// invokevirtual
// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	base_instructions.Index16Instruction
}

func (invokeVirtual *InvokeVirtual) Execute(frame *runtime_data_area.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	methodReference := constantPool.GetConstant(invokeVirtual.Index).(*heap.MethodReference)

	if methodReference.GetName() == "println" {
		operandStack := frame.GetOperandStack()

		switch methodReference.GetDescriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", operandStack.PopIntegerValue() != 0)
		case "(C)V":
			fmt.Printf("%c\n", operandStack.PopIntegerValue())
		case "(I)V", "(B)V", "(S)V":
			fmt.Printf("%v\n", operandStack.PopIntegerValue())
		case "(F)V":
			fmt.Printf("%v\n", operandStack.PopFloatValue())
		case "(J)V":
			fmt.Printf("%v\n", operandStack.PopLongValue())
		case "(D)V":
			fmt.Printf("%v\n", operandStack.PopDoubleValue())
		default:
			panic("println: " + methodReference.GetDescriptor())
		}

		operandStack.PopReferenceValue()
	}
}
