package base_instructions

import (
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

func InitializeClass(thread *runtime_data_area.Thread, class *heap.Class) {
	class.StartInitialization()
	executeClassInitializationMethod(thread, class)
}

func executeClassInitializationMethod(thread *runtime_data_area.Thread, class *heap.Class) {
	classInitializationMethod := class.GetClassInitializationMethod()

	if classInitializationMethod != nil {
		thread.PushFrame(thread.NewFrame(classInitializationMethod))
	}
}

func initializeSuperClass(thread *runtime_data_area.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.GetSuperClass()

		if superClass != nil && !superClass.IsInitializationStarted() {
			InitializeClass(thread, superClass)
		}
	}
}
