package heap

import (
	"strings"
)

type MethodDescriptorParser struct {
	descriptor       string
	offset           int
	methodDescriptor *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	methodDescriptorParser := &MethodDescriptorParser{}
	methodDescriptorParser.Parse(descriptor)

	return methodDescriptorParser.methodDescriptor
}

func (methodDescriptorParser *MethodDescriptorParser) Parse(descriptor string) {
	methodDescriptorParser.descriptor = descriptor
	methodDescriptorParser.methodDescriptor = &MethodDescriptor{}

	methodDescriptorParser.StartParseParameters()
	methodDescriptorParser.ParseParameterTypes()
	methodDescriptorParser.EndParseParameters()
	methodDescriptorParser.ParseReturnType()
	methodDescriptorParser.Finish()
}

func (methodDescriptorParser *MethodDescriptorParser) StartParseParameters() {
	if methodDescriptorParser.ReadUint8() != '(' {
		methodDescriptorParser.Panic()
	}
}

func (methodDescriptorParser *MethodDescriptorParser) EndParseParameters() {
	if methodDescriptorParser.ReadUint8() != ')' {
		methodDescriptorParser.Panic()
	}
}

func (methodDescriptorParser *MethodDescriptorParser) Finish() {
	if methodDescriptorParser.offset != len(methodDescriptorParser.descriptor) {
		methodDescriptorParser.Panic()
	}
}

func (methodDescriptorParser *MethodDescriptorParser) Panic() {
	panic("Bad method descriptor: " + methodDescriptorParser.descriptor)
}

func (methodDescriptorParser *MethodDescriptorParser) ReadUint8() uint8 {
	value := methodDescriptorParser.descriptor[methodDescriptorParser.offset]
	methodDescriptorParser.offset++

	return value
}

func (methodDescriptorParser *MethodDescriptorParser) UnReadUint8() {
	methodDescriptorParser.offset--
}

func (methodDescriptorParser *MethodDescriptorParser) ParseParameterTypes() {
	for {
		fieldType := methodDescriptorParser.ParseFieldType()

		if fieldType != "" {
			methodDescriptorParser.methodDescriptor.addParameterType(fieldType)
		} else {
			break
		}
	}
}

func (methodDescriptorParser *MethodDescriptorParser) ParseReturnType() {
	if methodDescriptorParser.ReadUint8() == 'V' {
		methodDescriptorParser.methodDescriptor.returnType = "V"

		return
	}

	methodDescriptorParser.UnReadUint8()

	fieldType := methodDescriptorParser.ParseFieldType()

	if fieldType != "" {
		methodDescriptorParser.methodDescriptor.returnType = fieldType

		return
	}

	methodDescriptorParser.Panic()
}

func (methodDescriptorParser *MethodDescriptorParser) ParseFieldType() string {
	switch methodDescriptorParser.ReadUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return methodDescriptorParser.ParseObjectType()
	case '[':
		return methodDescriptorParser.ParseArrayType()
	default:
		methodDescriptorParser.UnReadUint8()

		return ""
	}
}

func (methodDescriptorParser *MethodDescriptorParser) ParseObjectType() string {
	unreadDescriptor := methodDescriptorParser.descriptor[methodDescriptorParser.offset:]
	semicolonIndex := strings.IndexRune(unreadDescriptor, ';')

	if semicolonIndex == -1 {
		methodDescriptorParser.Panic()

		return ""
	}

	objectTypeStart := methodDescriptorParser.offset - 1
	objectTypeEnd := methodDescriptorParser.offset + semicolonIndex + 1
	methodDescriptorParser.offset = objectTypeEnd

	return methodDescriptorParser.descriptor[objectTypeStart:objectTypeEnd]
}

func (methodDescriptorParser *MethodDescriptorParser) ParseArrayType() string {
	arrayTypeStart := methodDescriptorParser.offset - 1

	methodDescriptorParser.ParseFieldType()

	arrayTypeEnd := methodDescriptorParser.offset

	return methodDescriptorParser.descriptor[arrayTypeStart:arrayTypeEnd]
}
