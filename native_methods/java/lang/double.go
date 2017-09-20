package lang

import (
	"math"

	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

const javaLangDouble = "java/lang/Double"

func init() {
	native_methods.RegisterNativeMethod(javaLangDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native_methods.RegisterNativeMethod(javaLangDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *runtime_data_area.Frame) {
	value := frame.GetLocalVariables().GetDoubleValue(0)
	bits := math.Float64bits(value)

	frame.GetOperandStack().PushLongValue(int64(bits))
}

func longBitsToDouble(frame *runtime_data_area.Frame) {
	bits := frame.GetLocalVariables().GetLongValue(0)
	value := math.Float64frombits(uint64(bits))

	frame.GetOperandStack().PushDoubleValue(value)
}
