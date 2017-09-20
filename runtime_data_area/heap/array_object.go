package heap

func (object *Object) GetByteArray() []int8 {
	return object.data.([]int8)
}

func (object *Object) GetShortArray() []int16 {
	return object.data.([]int16)
}

func (object *Object) GetIntArray() []int32 {
	return object.data.([]int32)
}

func (object *Object) GetLongArray() []int64 {
	return object.data.([]int64)
}

func (object *Object) GetCharArray() []uint16 {
	return object.data.([]uint16)
}

func (object *Object) GetFloatArray() []float32 {
	return object.data.([]float32)
}

func (object *Object) GetDoubleArray() []float64 {
	return object.data.([]float64)
}

func (object *Object) GetReferenceArray() []*Object {
	return object.data.([]*Object)
}

func (object *Object) GetArrayLength() int32 {
	switch object.data.(type) {
	case []int8:
		return int32(len(object.data.([]int8)))
	case []int16:
		return int32(len(object.data.([]int16)))
	case []int32:
		return int32(len(object.data.([]int32)))
	case []int64:
		return int32(len(object.data.([]int64)))
	case []uint16:
		return int32(len(object.data.([]uint16)))
	case []float32:
		return int32(len(object.data.([]float32)))
	case []float64:
		return int32(len(object.data.([]float64)))
	case []*Object:
		return int32(len(object.data.([]*Object)))
	default:
		panic("Invalid array type")
	}
}

func CopyArray(sourceArray, targetArray *Object, sourceArrayPosition, targetArrayPosition, length int32) {
	switch sourceArray.data.(type) {
	case []int8:
		copy(targetArray.data.([]int8)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]int8)[sourceArrayPosition:sourceArrayPosition+length])
	case []int16:
		copy(targetArray.data.([]int16)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]int16)[sourceArrayPosition:sourceArrayPosition+length])
	case []int32:
		copy(targetArray.data.([]int32)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]int32)[sourceArrayPosition:sourceArrayPosition+length])
	case []int64:
		copy(targetArray.data.([]int64)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]int64)[sourceArrayPosition:sourceArrayPosition+length])
	case []uint16:
		copy(targetArray.data.([]uint16)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]uint16)[sourceArrayPosition:sourceArrayPosition+length])
	case []float32:
		copy(targetArray.data.([]float32)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]float32)[sourceArrayPosition:sourceArrayPosition+length])
	case []float64:
		copy(targetArray.data.([]float64)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]float64)[sourceArrayPosition:sourceArrayPosition+length])
	case []*Object:
		copy(targetArray.data.([]*Object)[targetArrayPosition:targetArrayPosition+length], sourceArray.data.([]*Object)[sourceArrayPosition:sourceArrayPosition+length])
	default:
		panic("Not array!")
	}
}
