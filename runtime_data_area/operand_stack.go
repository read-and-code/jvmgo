package runtime_data_area

import "math"

type OperandStack struct {
	size     uint
	operands []Variable
}

func newOperandStack(maxStackSize uint) *OperandStack {
	if maxStackSize > 0 {
		return &OperandStack{
			operands: make([]Variable, maxStackSize),
		}
	}

	return nil
}

func (operandStack *OperandStack) PushIntegerValue(value int32) {
	operandStack.operands[operandStack.size].value = value
	operandStack.size++
}

func (operandStack *OperandStack) PopIntegerValue() int32 {
	operandStack.size--

	return operandStack.operands[operandStack.size].value
}

func (operandStack *OperandStack) PushFloatValue(value float32) {
	float32Bits := math.Float32bits(value)
	operandStack.operands[operandStack.size].value = int32(float32Bits)
	operandStack.size++
}

func (operandStack *OperandStack) PopFloatValue() float32 {
	operandStack.size--
	float32Bits := uint32(operandStack.operands[operandStack.size].value)

	return math.Float32frombits(float32Bits)
}

func (operandStack *OperandStack) PushLongValue(value int64) {
	operandStack.operands[operandStack.size].value = int32(value)
	operandStack.operands[operandStack.size+1].value = int32(value >> 32)
	operandStack.size += 2
}

func (operandStack *OperandStack) PopLongValue() int64 {
	operandStack.size -= 2
	low := uint32(operandStack.operands[operandStack.size].value)
	high := uint32(operandStack.operands[operandStack.size+1].value)

	return int64(high)<<32 | int64(low)
}

func (operandStack *OperandStack) PushDoubleValue(value float64) {
	float64Bits := math.Float64bits(value)
	operandStack.PushLongValue(int64(float64Bits))
}

func (operandStack *OperandStack) PopDoubleValue() float64 {
	float64Bits := uint64(operandStack.PopLongValue())

	return math.Float64frombits(float64Bits)
}

func (operandStack *OperandStack) PushReferenceValue(referenceValue *Object) {
	operandStack.operands[operandStack.size].referenceValue = referenceValue
	operandStack.size++
}

func (operandStack *OperandStack) PopReferenceValue() *Object {
	operandStack.size--
	referenceValue := operandStack.operands[operandStack.size].referenceValue
	operandStack.operands[operandStack.size].referenceValue = nil

	return referenceValue
}
