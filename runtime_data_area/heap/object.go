package heap

type Object struct {
	class *Class
	data  interface{}
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

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(object.class)
}
