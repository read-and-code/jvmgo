package runtime_data_area

type Frame struct {
	lower          *Frame
	localVariables LocalVariables
	operandStack   *OperandStack
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
