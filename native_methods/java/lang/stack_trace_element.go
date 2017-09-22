package lang

import "fmt"

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func (stackTraceElement *StackTraceElement) ToString() string {
	return fmt.Sprintf("%s.%s(%s:%d)", stackTraceElement.className, stackTraceElement.methodName, stackTraceElement.fileName, stackTraceElement.lineNumber)
}
