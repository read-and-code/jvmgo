package main

import (
	"fmt"
	"strings"

	"github.com/Frederick-S/jvmgo/classfile"
	"github.com/Frederick-S/jvmgo/classpath"
	"github.com/Frederick-S/jvmgo/runtime_data_area"
)

func main() {
	cmd := parseCmd()

	if cmd.showVersion {
		fmt.Println("Version 0.0.1")
	} else if cmd.showHelp || cmd.className == "" {
		printCmdUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	classloader := classpath.Parse(cmd.jrePath, cmd.classpath)
	className := strings.Replace(cmd.className, ".", "/", -1)
	classFile := loadClass(className, classloader)

	printClassInfo(classFile)

	frame := runtime_data_area.NewFrame(100, 100)

	testLocalVariables(frame.GetLocalVariables())
	testOperandStack(frame.GetOperandStack())
}

func loadClass(className string, classloader *classpath.Classloader) *classfile.ClassFile {
	classData, _, err := classloader.ReadClass(className)

	if err != nil {
		panic(err)
	}

	classFile, err := classfile.Parse(classData)

	if err != nil {
		panic(err)
	}

	return classFile
}

func printClassInfo(classFile *classfile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", classFile.GetMajorVersion(), classFile.GetMinorVersion())
	fmt.Printf("Constants count: %v\n", len(classFile.GetConstantPool()))
	fmt.Printf("Access flags: 0x%x\n", classFile.GetAccessFlags())
	fmt.Printf("This class: %v\n", classFile.GetClassName())
	fmt.Printf("Super class: %v\n", classFile.GetSuperClassName())
	fmt.Printf("Interfaces: %v\n", classFile.GetInterfaceNames())
	fmt.Printf("Fields count: %v\n", len(classFile.GetFields()))

	for _, field := range classFile.GetFields() {
		fmt.Printf(" %s\n", field.GetName())
	}

	fmt.Printf("Methods count: %v\n", len(classFile.GetMethods()))

	for _, method := range classFile.GetMethods() {
		fmt.Printf(" %s\n", method.GetName())
	}
}

func testLocalVariables(localVariables runtime_data_area.LocalVariables) {
	localVariables.SetIntegerValue(0, 100)
	localVariables.SetIntegerValue(1, -100)
	localVariables.SetLongValue(2, 2997924580)
	localVariables.SetLongValue(4, -2997924580)
	localVariables.SetFloatValue(6, 3.1415926)
	localVariables.SetDoubleValue(7, 2.71828182845)
	localVariables.SetReferenceValue(9, nil)

	println(localVariables.GetIntegerValue(0))
	println(localVariables.GetIntegerValue(1))
	println(localVariables.GetLongValue(2))
	println(localVariables.GetLongValue(4))
	println(localVariables.GetFloatValue(6))
	println(localVariables.GetDoubleValue(7))
	println(localVariables.GetReferenceValue(9))
}

func testOperandStack(operandStack *runtime_data_area.OperandStack) {
	operandStack.PushIntegerValue(100)
	operandStack.PushIntegerValue(-100)
	operandStack.PushLongValue(2997924580)
	operandStack.PushLongValue(-2997924580)
	operandStack.PushFloatValue(3.1415926)
	operandStack.PushDoubleValue(2.71828182845)
	operandStack.PushReferenceValue(nil)

	println(operandStack.PopReferenceValue())
	println(operandStack.PopDoubleValue())
	println(operandStack.PopFloatValue())
	println(operandStack.PopLongValue())
	println(operandStack.PopLongValue())
	println(operandStack.PopIntegerValue())
	println(operandStack.PopIntegerValue())
}
