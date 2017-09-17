package runtime_data_area

import "github.com/Frederick-S/jvmgo/runtime_data_area/heap"

type Thread struct {
	pc       int
	jvmStack *JVMStack
}

func NewThread() *Thread {
	return &Thread{
		jvmStack: newJVMStack(1024),
	}
}

func (thread *Thread) GetPC() int {
	return thread.pc
}

func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.jvmStack.PushFrame(frame)
}

func (thread *Thread) PopFrame() *Frame {
	return thread.jvmStack.PopFrame()
}

func (thread *Thread) GetCurrentFrame() *Frame {
	return thread.jvmStack.GetTopFrame()
}

func (thread *Thread) GetTopFrame() *Frame {
	return thread.jvmStack.GetTopFrame()
}

func (thread *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(thread, method)
}

func (thread *Thread) IsJVMStackEmpty() bool {
	return thread.jvmStack.IsEmpty()
}
