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
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchError(frame)

	loop(thread, method.GetCode())
}

func catchError(frame *runtime_data_area.Frame) {
	r := recover()

	if r != nil {
		// fmt.Printf("LocalVariables: %v\n", frame.GetLocalVariables())
		// fmt.Printf("OperandStack: %v\n", frame.GetOperandStack())

		// panic(r)
	}
}

func loop(thread *runtime_data_area.Thread, code []byte) {
	frame := thread.PopFrame()
	bytecodeReader := &base_instructions.BytecodeReader{}

	for {
		pc := frame.GetNextPC()

		thread.SetPC(pc)

		bytecodeReader.Reset(code, pc)
		operationCode := bytecodeReader.ReadUint8()
		instruction := instructions.NewInstruction(operationCode)
		instruction.FetchOperands(bytecodeReader)
		frame.SetNextPC(bytecodeReader.GetPC())

		fmt.Printf("PC: %2d instruction: %T %v\n", pc, instruction, instruction)
		instruction.Execute(frame)
	}
}
