package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getArrayClassName(className string) string {
	return "[" + convertClassNameToDescriptor(className)
}

func getArrayElementClassName(className string) string {
	if className[0] == '[' {
		descriptor := className[1:]

		return convertDescriptorToClassName(descriptor)
	}

	panic("Not array: " + className)
}

func convertClassNameToDescriptor(className string) string {
	if className[0] == '[' {
		// Array
		return className
	}

	descriptor, ok := primitiveTypes[className]

	if ok {
		// Primitive
		return descriptor
	}

	// Object
	return "L" + className + ";"
}

func convertDescriptorToClassName(descriptor string) string {
	if descriptor[0] == '[' {
		// Array
		return descriptor
	}

	if descriptor[0] == 'L' {
		// Object
		return descriptor[1 : len(descriptor)-1]
	}

	for primitiveType, primitiveTypeDescriptor := range primitiveTypes {
		if primitiveTypeDescriptor == descriptor {
			// Primitive
			return primitiveType
		}
	}

	panic("Invalid descriptor: " + descriptor)
}
