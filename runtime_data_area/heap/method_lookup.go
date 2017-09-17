package heap

func lookupMethodInClass(class *Class, name, descriptor string) *Method {
	for currentClass := class; currentClass != nil; currentClass = currentClass.superClass {
		for _, method := range currentClass.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}

	return nil
}

func lookupMethodInInterfaces(interfaces []*Class, name, descriptor string) *Method {
	for _, currentInterface := range interfaces {
		for _, method := range currentInterface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}

		method := lookupMethodInInterfaces(currentInterface.interfaces, name, descriptor)

		if method != nil {
			return method
		}
	}

	return nil
}
