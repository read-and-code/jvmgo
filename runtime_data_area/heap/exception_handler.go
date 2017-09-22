package heap

type ExceptionHandler struct {
	startPC   int
	endPC     int
	handlerPC int
	catchType *ClassReference
}
