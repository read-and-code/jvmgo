package runtime_data_area

type Frame struct {
	lower          *Frame
	localVariables LocalVariables
	operandStack   *OperandStack
	thread         *Thread
	nextPC         int
}

func NewFrame(maxNumberOfLocalVariables, maxStackSize uint) *Frame {
	return &Frame{
		localVariables: newLocalVariables(maxNumberOfLocalVariables),
		operandStack:   newOperandStack(maxStackSize),
	}
}

func (frame *Frame) GetLocalVariables() LocalVariables {
	return frame.localVariables
}

func (frame *Frame) GetOperandStack() *OperandStack {
	return frame.operandStack
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
