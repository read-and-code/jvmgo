package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (methodDescriptor *MethodDescriptor) addParameterType(parameterType string) {
	parameterTypesCount := len(methodDescriptor.parameterTypes)

	if parameterTypesCount == cap(methodDescriptor.parameterTypes) {
		parameterTypes := make([]string, parameterTypesCount, parameterTypesCount+4)

		copy(parameterTypes, methodDescriptor.parameterTypes)

		methodDescriptor.parameterTypes = parameterTypes
	}

	methodDescriptor.parameterTypes = append(methodDescriptor.parameterTypes, parameterType)
}
