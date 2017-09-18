package heap

func (class *Class) IsArray() bool {
	return class.name[0] == '['
}

func (class *Class) GetArrayElementClass() *Class {
	className := getArrayElementClassName(class.name)

	return class.classLoader.LoadClass(className)
}

func (class *Class) NewArray(length uint) *Object {
	if !class.IsArray() {
		panic("Not array class: " + class.name)
	}

	switch class.name {
	case "[Z":
		return &Object{class, make([]int8, length)}
	case "[B":
		return &Object{class, make([]int8, length)}
	case "[C":
		return &Object{class, make([]uint16, length)}
	case "[S":
		return &Object{class, make([]int16, length)}
	case "[I":
		return &Object{class, make([]int32, length)}
	case "[J":
		return &Object{class, make([]int64, length)}
	case "[F":
		return &Object{class, make([]float32, length)}
	case "[D":
		return &Object{class, make([]float64, length)}
	default:
		return &Object{class, make([]*Object, length)}
	}
}
