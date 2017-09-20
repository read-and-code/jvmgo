package native_methods

var nativeMethodsRegistry = map[string]NativeMethod{}

func RegisterNativeMethod(className, methodName, methodDescriptor string, nativeMethod NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor

	nativeMethodsRegistry[key] = nativeMethod
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor

	nativeMethod, ok := nativeMethodsRegistry[key]

	if ok {
		return nativeMethod
	}

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	return nil
}
