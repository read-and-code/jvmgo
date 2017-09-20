package heap

type Object struct {
	class     *Class
	data      interface{}
	extraData interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newVariables(class.instanceVariablesCount),
	}
}

func (object *Object) GetClass() *Class {
	return object.class
}

func (object *Object) GetFields() Variables {
	return object.data.(Variables)
}

func (object *Object) GetExtraData() interface{} {
	return object.extraData
}

func (object *Object) SetExtraData(extraData interface{}) {
	object.extraData = extraData
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(object.class)
}

func (object *Object) GetReferenceValue(name, descriptor string) *Object {
	field := object.class.GetField(name, descriptor, false)

	return object.data.(Variables).GetReferenceValue(field.variableIndex)
}

func (object *Object) SetReferenceValue(name, descriptor string, referenceValue *Object) {
	field := object.class.GetField(name, descriptor, false)

	object.data.(Variables).SetReferenceValue(field.variableIndex, referenceValue)
}

func (object *Object) Clone() *Object {
	return &Object{
		class: object.class,
		data:  object.clone(),
	}
}

func (object *Object) clone() interface{} {
	switch object.data.(type) {
	case []int8:
		data := object.data.([]int8)
		clonedData := make([]int8, len(data))

		copy(clonedData, data)

		return clonedData
	case []int16:
		data := object.data.([]int16)
		clonedData := make([]int16, len(data))

		copy(clonedData, data)

		return clonedData
	case []uint16:
		data := object.data.([]uint16)
		clonedData := make([]uint16, len(data))

		copy(clonedData, data)

		return clonedData
	case []int32:
		data := object.data.([]int32)
		clonedData := make([]int32, len(data))

		copy(clonedData, data)

		return clonedData
	case []int64:
		data := object.data.([]int64)
		clonedData := make([]int64, len(data))

		copy(clonedData, data)

		return clonedData
	case []float32:
		data := object.data.([]float32)
		clonedData := make([]float32, len(data))

		copy(clonedData, data)

		return clonedData
	case []float64:
		data := object.data.([]float64)
		clonedData := make([]float64, len(data))

		copy(clonedData, data)

		return clonedData
	case []*Object:
		data := object.data.([]*Object)
		clonedData := make([]*Object, len(data))

		copy(clonedData, data)

		return clonedData
	default:
		data := object.data.(Variables)
		clonedData := newVariables(uint(len(data)))

		copy(clonedData, data)

		return clonedData
	}
}
