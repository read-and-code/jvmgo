package runtime_data_area

import (
	"math"

	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

type LocalVariables []Variable

func newLocalVariables(maxNumberOfLocalVariables uint) LocalVariables {
	if maxNumberOfLocalVariables > 0 {
		return make([]Variable, maxNumberOfLocalVariables)
	}

	return nil
}

func (localVariables LocalVariables) SetIntegerValue(index uint, value int32) {
	localVariables[index].numericalValue = value
}

func (localVariables LocalVariables) GetIntegerValue(index uint) int32 {
	return localVariables[index].numericalValue
}

func (localVariables LocalVariables) SetFloatValue(index uint, value float32) {
	float32Bits := math.Float32bits(value)
	localVariables[index].numericalValue = int32(float32Bits)
}

func (localVariables LocalVariables) GetFloatValue(index uint) float32 {
	float32Bits := uint32(localVariables[index].numericalValue)

	return math.Float32frombits(float32Bits)
}

func (localVariables LocalVariables) SetLongValue(index uint, value int64) {
	localVariables[index].numericalValue = int32(value)
	localVariables[index+1].numericalValue = int32(value >> 32)
}

func (localVariables LocalVariables) GetLongValue(index uint) int64 {
	low := uint32(localVariables[index].numericalValue)
	high := uint32(localVariables[index+1].numericalValue)

	return int64(high)<<32 | int64(low)
}

func (localVariables LocalVariables) SetDoubleValue(index uint, value float64) {
	float64Bits := math.Float64bits(value)
	localVariables.SetLongValue(index, int64(float64Bits))
}

func (localVariables LocalVariables) GetDoubleValue(index uint) float64 {
	float64Bits := uint64(localVariables.GetLongValue(index))

	return math.Float64frombits(float64Bits)
}

func (localVariables LocalVariables) SetReferenceValue(index uint, value *heap.Object) {
	localVariables[index].referenceValue = value
}

func (localVariables LocalVariables) GetReferenceValue(index uint) *heap.Object {
	return localVariables[index].referenceValue
}

func (localVariables LocalVariables) SetVariable(index uint, variable Variable) {
	localVariables[index] = variable
}
