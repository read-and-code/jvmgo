package heap

import "math"

type Variables []Variable

func newVariables(variableCount uint) Variables {
	if variableCount > 0 {
		return make([]Variable, variableCount)
	}

	return nil
}

func (variables Variables) SetIntegerValue(index uint, value int32) {
	variables[index].numericalValue = value
}

func (variables Variables) GetIntegerValue(index uint) int32 {
	return variables[index].numericalValue
}

func (variables Variables) SetFloatValue(index uint, value float32) {
	float32Bits := math.Float32bits(value)
	variables[index].numericalValue = int32(float32Bits)
}

func (variables Variables) GetFloatValue(index uint) float32 {
	float32Bits := uint32(variables[index].numericalValue)

	return math.Float32frombits(float32Bits)
}

func (variables Variables) SetLongValue(index uint, value int64) {
	variables[index].numericalValue = int32(value)
	variables[index+1].numericalValue = int32(value >> 32)
}

func (variables Variables) GetLongValue(index uint) int64 {
	low := uint32(variables[index].numericalValue)
	high := uint32(variables[index+1].numericalValue)

	return int64(high)<<32 | int64(low)
}

func (variables Variables) SetDoubleValue(index uint, value float64) {
	float64Bits := math.Float64bits(value)
	variables.SetLongValue(index, int64(float64Bits))
}

func (variables Variables) GetDoubleValue(index uint) float64 {
	float64Bits := uint64(variables.GetLongValue(index))

	return math.Float64frombits(float64Bits)
}

func (variables Variables) SetReferenceValue(index uint, value *Object) {
	variables[index].referenceValue = value
}

func (variables Variables) GetReferenceValue(index uint) *Object {
	return variables[index].referenceValue
}
