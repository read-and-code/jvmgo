package lang

import (
	"math"

	"github.com/Frederick-S/jvmgo/native_methods"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

const javaLangFloat = "java/lang/Float"

func init() {
	native_methods.RegisterNativeMethod(javaLangFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native_methods.RegisterNativeMethod(javaLangFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}

func floatToRawIntBits(frame *runtime_data_area.Frame) {
	value := frame.GetLocalVariables().GetFloatValue(0)
	bits := math.Float32bits(value)

	frame.GetOperandStack().PushIntegerValue(int32(bits))
}

func intBitsToFloat(frame *runtime_data_area.Frame) {
	bits := frame.GetLocalVariables().GetIntegerValue(0)
	value := math.Float32frombits(uint32(bits))

	frame.GetOperandStack().PushFloatValue(value)
}
