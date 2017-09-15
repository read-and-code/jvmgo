package main

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/classfile"
	"github.com/Frederick-S/jvmgo/instructions"
	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttribute := methodInfo.GetCodeAttribute()
	thread := runtime_data_area.NewThread()
	frame := thread.NewFrame(codeAttribute.GetMaxNumberOfLocalVariables(), codeAttribute.GetMaxStackSize())

	thread.PushFrame(frame)

	defer catchError(frame)

	loop(thread, codeAttribute.GetCode())
}

func catchError(frame *runtime_data_area.Frame) {
	r := recover()

	if r != nil {
		fmt.Printf("LocalVariables: %v\n", frame.GetLocalVariables())
		fmt.Printf("OperandStack: %v\n", frame.GetOperandStack())

		panic(r)
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
