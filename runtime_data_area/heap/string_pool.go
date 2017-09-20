package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

func ConvertGoStringToJavaString(classLoader *ClassLoader, goString string) *Object {
	internedString, ok := internedStrings[goString]

	if ok {
		return internedString
	}

	charArray := convertStringToUtf16(goString)
	javaCharArray := &Object{classLoader.LoadClass("[C"), charArray, nil}

	javaString := classLoader.LoadClass("java/lang/String").NewObject()
	javaString.SetReferenceValue("value", "[C", javaCharArray)

	internedStrings[goString] = javaString

	return javaString
}

func ConvertJavaStringToGoString(javaString *Object) string {
	return convertUtf16ToString(javaString.GetReferenceValue("value", "[C").GetCharArray())
}

func convertStringToUtf16(value string) []uint16 {
	return utf16.Encode([]rune(value))
}

func convertUtf16ToString(chars []uint16) string {
	return string(utf16.Decode(chars))
}

func GetInternedString(javaString *Object) *Object {
	goString := ConvertJavaStringToGoString(javaString)

	internedString, ok := internedStrings[goString]

	if ok {
		return internedString
	}

	internedStrings[goString] = javaString

	return javaString
}
