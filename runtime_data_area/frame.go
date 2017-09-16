package runtime_data_area

import "github.com/Frederick-S/jvmgo/runtime_data_area/heap"

type Frame struct {
	lower          *Frame
	localVariables LocalVariables
	operandStack   *OperandStack
	thread         *Thread
	method         *heap.Method
	nextPC         int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:         thread,
		method:         method,
		localVariables: newLocalVariables(method.GetMaxNumberOfLocalVariables()),
		operandStack:   newOperandStack(method.GetMaxStackSize()),
	}
}

func (frame *Frame) GetLocalVariables() LocalVariables {
	return frame.localVariables
}

func (frame *Frame) GetOperandStack() *OperandStack {
	return frame.operandStack
}

func (frame *Frame) GetMethod() *heap.Method {
	return frame.method
}

func (frame *Frame) GetThread() *Thread {
	return frame.thread
}

func (frame *Frame) GetNextPC() int {
	return frame.nextPC
}

func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}
