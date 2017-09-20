package native_methods

import "github.com/Frederick-S/jvmgo/runtime_data_area"

type NativeMethod func(frame *runtime_data_area.Frame)

func emptyNativeMethod(frame *runtime_data_area.Frame) {
}
