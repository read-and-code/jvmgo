package instructions

import (
	"fmt"

	"github.com/Frederick-S/jvmgo/instructions/base_instructions"
	"github.com/Frederick-S/jvmgo/instructions/comparison_instructions"
	"github.com/Frederick-S/jvmgo/instructions/constant_instructions"
	"github.com/Frederick-S/jvmgo/instructions/control_instructions"
	"github.com/Frederick-S/jvmgo/instructions/conversion_instructions"
	"github.com/Frederick-S/jvmgo/instructions/extended_instructions"
	"github.com/Frederick-S/jvmgo/instructions/load_instructions"
	"github.com/Frederick-S/jvmgo/instructions/math_instructions"
	"github.com/Frederick-S/jvmgo/instructions/reference_instructions"
	"github.com/Frederick-S/jvmgo/instructions/stack_instructions"
	"github.com/Frederick-S/jvmgo/instructions/store_instructions"
)

var (
	nop         = &constant_instructions.Nop{}
	aconst_null = &constant_instructions.AConstNull{}
	iconst_m1   = &constant_instructions.IConstM1{}
	iconst_0    = &constant_instructions.IConst0{}
	iconst_1    = &constant_instructions.IConst1{}
	iconst_2    = &constant_instructions.IConst2{}
	iconst_3    = &constant_instructions.IConst3{}
	iconst_4    = &constant_instructions.IConst4{}
	iconst_5    = &constant_instructions.IConst5{}
	lconst_0    = &constant_instructions.LConst0{}
	lconst_1    = &constant_instructions.LConst1{}
	fconst_0    = &constant_instructions.FConst0{}
	fconst_1    = &constant_instructions.FConst1{}
	fconst_2    = &constant_instructions.FConst2{}
	dconst_0    = &constant_instructions.DConst0{}
	dconst_1    = &constant_instructions.DConst1{}
	iload_0     = &load_instructions.ILoad0{}
	iload_1     = &load_instructions.ILoad1{}
	iload_2     = &load_instructions.ILoad2{}
	iload_3     = &load_instructions.ILoad3{}
	lload_0     = &load_instructions.LLoad0{}
	lload_1     = &load_instructions.LLoad1{}
	lload_2     = &load_instructions.LLoad2{}
	lload_3     = &load_instructions.LLoad3{}
	fload_0     = &load_instructions.FLoad0{}
	fload_1     = &load_instructions.FLoad1{}
	fload_2     = &load_instructions.FLoad2{}
	fload_3     = &load_instructions.FLoad3{}
	dload_0     = &load_instructions.DLoad0{}
	dload_1     = &load_instructions.DLoad1{}
	dload_2     = &load_instructions.DLoad2{}
	dload_3     = &load_instructions.DLoad3{}
	aload_0     = &load_instructions.ALoad0{}
	aload_1     = &load_instructions.ALoad1{}
	aload_2     = &load_instructions.ALoad2{}
	aload_3     = &load_instructions.ALoad3{}
	// iaload      = &IALOAD{}
	// laload      = &LALOAD{}
	// faload      = &FALOAD{}
	// daload      = &DALOAD{}
	// aaload      = &AALOAD{}
	// baload      = &BALOAD{}
	// caload      = &CALOAD{}
	// saload      = &SALOAD{}
	istore_0 = &store_instructions.IStore0{}
	istore_1 = &store_instructions.IStore1{}
	istore_2 = &store_instructions.IStore2{}
	istore_3 = &store_instructions.IStore3{}
	lstore_0 = &store_instructions.LStore0{}
	lstore_1 = &store_instructions.LStore1{}
	lstore_2 = &store_instructions.LStore2{}
	lstore_3 = &store_instructions.LStore3{}
	fstore_0 = &store_instructions.FStore0{}
	fstore_1 = &store_instructions.FStore1{}
	fstore_2 = &store_instructions.FStore2{}
	fstore_3 = &store_instructions.FStore3{}
	dstore_0 = &store_instructions.DStore0{}
	dstore_1 = &store_instructions.DStore1{}
	dstore_2 = &store_instructions.DStore2{}
	dstore_3 = &store_instructions.DStore3{}
	astore_0 = &store_instructions.AStore0{}
	astore_1 = &store_instructions.AStore1{}
	astore_2 = &store_instructions.AStore2{}
	astore_3 = &store_instructions.AStore3{}
	// iastore  = &IASTORE{}
	// lastore  = &LASTORE{}
	// fastore  = &FASTORE{}
	// dastore  = &DASTORE{}
	// aastore  = &AASTORE{}
	// bastore  = &BASTORE{}
	// castore  = &CASTORE{}
	// sastore  = &SASTORE{}
	pop     = &stack_instructions.Pop{}
	pop2    = &stack_instructions.Pop2{}
	dup     = &stack_instructions.Dup{}
	dup_x1  = &stack_instructions.DupX1{}
	dup_x2  = &stack_instructions.DupX2{}
	dup2    = &stack_instructions.Dup2{}
	dup2_x1 = &stack_instructions.Dup2X1{}
	dup2_x2 = &stack_instructions.Dup2X2{}
	swap    = &stack_instructions.Swap{}
	iadd    = &math_instructions.IAdd{}
	ladd    = &math_instructions.LAdd{}
	fadd    = &math_instructions.FAdd{}
	dadd    = &math_instructions.DAdd{}
	isub    = &math_instructions.ISub{}
	lsub    = &math_instructions.LSub{}
	fsub    = &math_instructions.FSub{}
	dsub    = &math_instructions.DSub{}
	imul    = &math_instructions.IMul{}
	lmul    = &math_instructions.LMul{}
	fmul    = &math_instructions.FMul{}
	dmul    = &math_instructions.DMul{}
	idiv    = &math_instructions.IDiv{}
	ldiv    = &math_instructions.LDiv{}
	fdiv    = &math_instructions.FDiv{}
	ddiv    = &math_instructions.DDiv{}
	irem    = &math_instructions.IRem{}
	lrem    = &math_instructions.LRem{}
	frem    = &math_instructions.FRem{}
	drem    = &math_instructions.DRem{}
	ineg    = &math_instructions.INeg{}
	lneg    = &math_instructions.LNeg{}
	fneg    = &math_instructions.FNeg{}
	dneg    = &math_instructions.DNeg{}
	ishl    = &math_instructions.IShl{}
	lshl    = &math_instructions.LShl{}
	ishr    = &math_instructions.IShr{}
	lshr    = &math_instructions.LShr{}
	iushr   = &math_instructions.IUshr{}
	lushr   = &math_instructions.LUshr{}
	iand    = &math_instructions.IAnd{}
	land    = &math_instructions.LAnd{}
	ior     = &math_instructions.IOr{}
	lor     = &math_instructions.LOr{}
	ixor    = &math_instructions.IXor{}
	lxor    = &math_instructions.LXor{}
	i2l     = &conversion_instructions.I2l{}
	i2f     = &conversion_instructions.I2f{}
	i2d     = &conversion_instructions.I2d{}
	l2i     = &conversion_instructions.L2i{}
	l2f     = &conversion_instructions.L2f{}
	l2d     = &conversion_instructions.L2d{}
	f2i     = &conversion_instructions.F2i{}
	f2l     = &conversion_instructions.F2l{}
	f2d     = &conversion_instructions.F2d{}
	d2i     = &conversion_instructions.D2i{}
	d2l     = &conversion_instructions.D2l{}
	d2f     = &conversion_instructions.D2f{}
	i2b     = &conversion_instructions.I2b{}
	i2c     = &conversion_instructions.I2c{}
	i2s     = &conversion_instructions.I2s{}
	lcmp    = &comparison_instructions.LCmp{}
	fcmpl   = &comparison_instructions.FCmpl{}
	fcmpg   = &comparison_instructions.FCmpg{}
	dcmpl   = &comparison_instructions.DCmpl{}
	dcmpg   = &comparison_instructions.DCmpg{}
	// ireturn = &IRETURN{}
	// lreturn = &LRETURN{}
	// freturn = &FRETURN{}
	// dreturn = &DRETURN{}
	// areturn = &ARETURN{}
	// _return = &RETURN{}
	// arraylength   = &ARRAY_LENGTH{}
	// athrow        = &ATHROW{}
	// monitorenter  = &MONITOR_ENTER{}
	// monitorexit   = &MONITOR_EXIT{}
	// invoke_native = &INVOKE_NATIVE{}
)

func NewInstruction(operationCode byte) base_instructions.Instruction {
	switch operationCode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return &constant_instructions.BIPush{}
	case 0x11:
		return &constant_instructions.SIPush{}
	case 0x12:
		return &constant_instructions.Ldc{}
	case 0x13:
		return &constant_instructions.LdcW{}
	case 0x14:
		return &constant_instructions.Ldc2W{}
	case 0x15:
		return &load_instructions.ILoad{}
	case 0x16:
		return &load_instructions.LLoad{}
	case 0x17:
		return &load_instructions.FLoad{}
	case 0x18:
		return &load_instructions.DLoad{}
	case 0x19:
		return &load_instructions.ALoad{}
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	// case 0x2e:
	// 	return iaload
	// case 0x2f:
	// 	return laload
	// case 0x30:
	// 	return faload
	// case 0x31:
	// 	return daload
	// case 0x32:
	// 	return aaload
	// case 0x33:
	// 	return baload
	// case 0x34:
	// 	return caload
	// case 0x35:
	// 	return saload
	case 0x36:
		return &store_instructions.IStore{}
	case 0x37:
		return &store_instructions.LStore{}
	case 0x38:
		return &store_instructions.FStore{}
	case 0x39:
		return &store_instructions.DStore{}
	case 0x3a:
		return &store_instructions.AStore{}
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	// case 0x4f:
	// 	return iastore
	// case 0x50:
	// 	return lastore
	// case 0x51:
	// 	return fastore
	// case 0x52:
	// 	return dastore
	// case 0x53:
	// 	return aastore
	// case 0x54:
	// 	return bastore
	// case 0x55:
	// 	return castore
	// case 0x56:
	// 	return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &math_instructions.IInc{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return &comparison_instructions.IfEq{}
	case 0x9a:
		return &comparison_instructions.IfNe{}
	case 0x9b:
		return &comparison_instructions.IfLt{}
	case 0x9c:
		return &comparison_instructions.IfGe{}
	case 0x9d:
		return &comparison_instructions.IfGt{}
	case 0x9e:
		return &comparison_instructions.IfLe{}
	case 0x9f:
		return &comparison_instructions.IfICmpEq{}
	case 0xa0:
		return &comparison_instructions.IfICmpNe{}
	case 0xa1:
		return &comparison_instructions.IfICmpLt{}
	case 0xa2:
		return &comparison_instructions.IfICmpGe{}
	case 0xa3:
		return &comparison_instructions.IfICmpGt{}
	case 0xa4:
		return &comparison_instructions.IfICmpLe{}
	case 0xa5:
		return &comparison_instructions.IfACmpEq{}
	case 0xa6:
		return &comparison_instructions.IfACmpNe{}
	case 0xa7:
		return &control_instructions.GoTo{}
	// case 0xa8:
	// 	return &JSR{}
	// case 0xa9:
	// 	return &RET{}
	case 0xaa:
		return &control_instructions.TableSwitch{}
	case 0xab:
		return &control_instructions.LookupSwitch{}
		// case 0xac:
		// 	return ireturn
		// case 0xad:
		// 	return lreturn
		// case 0xae:
		// 	return freturn
		// case 0xaf:
		// 	return dreturn
		// case 0xb0:
		// 	return areturn
		// case 0xb1:
		// 	return _return
	case 0xb2:
		return &reference_instructions.GetStatic{}
	case 0xb3:
		return &reference_instructions.PutStatic{}
	case 0xb4:
		return &reference_instructions.GetField{}
	case 0xb5:
		return &reference_instructions.PutField{}
	case 0xb6:
		return &reference_instructions.InvokeVirtual{}
	case 0xb7:
		return &reference_instructions.InvokeSpecial{}
		// case 0xb8:
		// 	return &INVOKE_STATIC{}
		// case 0xb9:
		// 	return &INVOKE_INTERFACE{}
		// case 0xba:
		// 	return &INVOKE_DYNAMIC{}
	case 0xbb:
		return &reference_instructions.New{}
		// case 0xbc:
		// 	return &NEW_ARRAY{}
		// case 0xbd:
		// 	return &ANEW_ARRAY{}
		// case 0xbe:
		// 	return arraylength
		// case 0xbf:
		// 	return athrow
	case 0xc0:
		return &reference_instructions.CheckCast{}
	case 0xc1:
		return &reference_instructions.InstanceOf{}
	// case 0xc2:
	// 	return monitorenter
	// case 0xc3:
	// 	return monitorexit
	case 0xc4:
		return &extended_instructions.Wide{}
	// case 0xc5:
	// 	return &MULTI_ANEW_ARRAY{}
	case 0xc6:
		return &extended_instructions.IfNull{}
	case 0xc7:
		return &extended_instructions.IfNoNull{}
	case 0xc8:
		return &extended_instructions.GoToW{}
	// case 0xc9:
	// 	return &JSR_W{}
	// case 0xca: breakpoint
	// case 0xfe: impdep1
	// case 0xff: impdep2
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", operationCode))
	}
}
