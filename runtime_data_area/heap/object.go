package heap

type Object struct {
	class  *Class
	fields Variables
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newVariables(class.instanceVariablesCount),
	}
}

func (object *Object) GetClass() *Class {
	return object.class
}

func (object *Object) GetFields() Variables {
	return object.fields
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(object.class)
}
