package base_instructions

import "github.com/Frederick-S/jvmgo/runtime_data_area"

func JumpToBranch(frame *runtime_data_area.Frame, offset int) {
	pc := frame.GetThread().GetPC()
	nextPC := pc + offset

	frame.SetNextPC(nextPC)
}
