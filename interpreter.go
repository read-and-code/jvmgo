package main

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/instructions"
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
	"github.com/Frederick-S/jvmgo/runtime_data_area/heap"
)

func interpret(method *heap.Method) {
	thread := runtime_data_area.NewThread()
	thread.PushFrame(thread.NewFrame(method))

	defer catchError(thread)

	loop(thread)
}

func catchError(thread *runtime_data_area.Thread) {
	r := recover()

	if r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *runtime_data_area.Thread) {
	bytecodeReader := &base_instructions.BytecodeReader{}

	for {
		frame := thread.GetCurrentFrame()
		pc := frame.GetNextPC()

		thread.SetPC(pc)

		bytecodeReader.Reset(frame.GetMethod().GetCode(), pc)
		operationCode := bytecodeReader.ReadUint8()
		instruction := instructions.NewInstruction(operationCode)
		instruction.FetchOperands(bytecodeReader)
		frame.SetNextPC(bytecodeReader.GetPC())

		logInstruction(frame, instruction)
		instruction.Execute(frame)

		if thread.IsJVMStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *runtime_data_area.Frame, instruction base_instructions.Instruction) {
	method := frame.GetMethod()
	className := method.GetClass().GetName()
	methodName := method.GetName()
	pc := frame.GetThread().GetPC()

	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, instruction, instruction)
}

func logFrames(thread *runtime_data_area.Thread) {
	for !thread.IsJVMStackEmpty() {
		frame := thread.PopFrame()
		method := frame.GetMethod()
		className := method.GetClass().GetName()

		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.GetNextPC(), className, method.GetName(), method.GetDescriptor())
	}
}
