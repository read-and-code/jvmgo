package heap

import "github.com/Frederick-S/jvmgo/classfile"

type ExceptionTable []*ExceptionHandler

func newExceptionTable(exceptionTableEntries []*classfile.ExceptionTableEntry, constantPool *ConstantPool) ExceptionTable {
	exceptionTable := make([]*ExceptionHandler, len(exceptionTableEntries))

	for i, exceptionTableEntry := range exceptionTableEntries {
		exceptionTable[i] = &ExceptionHandler{
			startPC:   int(exceptionTableEntry.GetStartPC()),
			endPC:     int(exceptionTableEntry.GetEndPC()),
			handlerPC: int(exceptionTableEntry.GetHandlerPC()),
			catchType: getCatchType(uint(exceptionTableEntry.GetCatchTypeIndex()), constantPool),
		}
	}

	return exceptionTable
}

func getCatchType(index uint, constantPool *ConstantPool) *ClassReference {
	if index == 0 {
		return nil
	}

	return constantPool.GetConstant(index).(*ClassReference)
}

func (exceptionTable ExceptionTable) findExceptionHandler(exceptionClass *Class, pc int) *ExceptionHandler {
	for _, exceptionHandler := range exceptionTable {
		if pc >= exceptionHandler.startPC && pc < exceptionHandler.endPC {
			if exceptionHandler.catchType == nil {
				return exceptionHandler
			}

			catchClass := exceptionHandler.catchType.GetResolvedClass()

			if catchClass == exceptionClass || catchClass.IsSuperClassOf(exceptionClass) {
				return exceptionHandler
			}
		}
	}

	return nil
}
