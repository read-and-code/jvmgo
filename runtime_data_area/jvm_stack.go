package runtime_data_area

type JVMStack struct {
	maxSize  uint
	size     uint
	topFrame *Frame
}

func newJVMStack(maxStackSize uint) *JVMStack {
	return &JVMStack{
		maxSize: maxStackSize,
	}
}

func (jvmStack *JVMStack) PushFrame(frame *Frame) {
	if jvmStack.size >= jvmStack.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if jvmStack.topFrame != nil {
		frame.lower = jvmStack.topFrame
	}

	jvmStack.topFrame = frame
	jvmStack.size++
}

func (jvmStack *JVMStack) PopFrame() *Frame {
	if jvmStack.topFrame == nil {
		panic("JVM stack is empty")
	}

	topFrame := jvmStack.topFrame
	jvmStack.topFrame = topFrame.lower
	topFrame.lower = nil
	jvmStack.size--

	return topFrame
}

func (jvmStack *JVMStack) GetTopFrame() *Frame {
	if jvmStack.topFrame == nil {
		panic("JVM stack is empty")
	}

	return jvmStack.topFrame
}
