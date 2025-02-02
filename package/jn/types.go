package jn

const (
	Void uint8 = 0
	I8   uint8 = 1
	I16  uint8 = 2
	I32  uint8 = 3
	I64  uint8 = 4
	U8   uint8 = 5
	U16  uint8 = 6
	U32  uint8 = 7
	U64  uint8 = 8
	Bool uint8 = 9
	Str  uint8 = 10
	F32  uint8 = 11
	F64  uint8 = 12
	Any  uint8 = 13
	Rune uint8 = 14
	Id   uint8 = 15
	Func uint8 = 16
	Nil  uint8 = 17
	Size uint8 = 18
	Map  uint8 = 19
)

const (
	NumericTypeStr = "<numeric>"
	NilTypeStr     = "<nil>"
)

func TypeGreaterThan(t1, t2 uint8) bool {
	switch t1 {
	case I16:
		return t2 == I8
	case I32:
		return t2 == I8 ||
			t2 == I16
	case I64:
		return t2 == I8 ||
			t2 == I16 ||
			t2 == I32
	case U16:
		return t2 == U8
	case U32:
		return t2 == U8 ||
			t2 == U16
	case U64:
		return t2 == U8 ||
			t2 == U16 ||
			t2 == U32
	case F32:
		return t2 != Any && t2 != F64
	case F64:
		return t2 != Any
	case Size:
		return true
	}
	return false
}

func TypesAreCompatible(t1, t2 uint8, ignoreany bool) bool {
	if !ignoreany && t2 == Any {
		return true
	}
	switch t1 {
	case Any:
		if ignoreany {
			return false
		}
		return true
	case I8:
		return t2 == I8 ||
			t2 == I16 ||
			t2 == I32 ||
			t2 == I64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case I16:
		return t2 == I16 ||
			t2 == I32 ||
			t2 == I64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case I32:
		return t2 == I32 ||
			t2 == I64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case I64:
		return t2 == I64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case U8:
		return t2 == U8 ||
			t2 == U16 ||
			t2 == U32 ||
			t2 == U64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case U16:
		return t2 == U16 ||
			t2 == U32 ||
			t2 == U64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case U32:
		return t2 == U32 ||
			t2 == U64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case U64, Size:
		return t2 == U64 ||
			t2 == F32 ||
			t2 == F64 ||
			t2 == Size
	case Bool:
		return t2 == Bool
	case Str:
		return t2 == Str
	case F32:
		return t2 == F32 ||
			t2 == F64
	case F64:
		return t2 == F64
	case Rune:
		return t2 == Rune ||
			t2 == I32 ||
			t2 == I64 ||
			t2 == U16 ||
			t2 == U32 ||
			t2 == U64 ||
			t2 == Size
	case Nil:
		return t2 == Nil
	}
	return false
}

func IsIntegerType(t uint8) bool {
	return IsSignedIntegerType(t) || IsUnsignedNumericType(t)
}

func IsNumericType(t uint8) bool {
	return IsIntegerType(t) || IsFloatType(t)
}

func IsFloatType(t uint8) bool {
	return t == F32 || t == F64
}

func IsSignedNumericType(t uint8) bool {
	return IsSignedIntegerType(t) ||
		t == F32 ||
		t == F64
}

func IsSignedIntegerType(t uint8) bool {
	return t == I8 ||
		t == I16 ||
		t == I32 ||
		t == I64
}

func IsUnsignedNumericType(t uint8) bool {
	return t == U8 ||
		t == U16 ||
		t == U32 ||
		t == U64 ||
		t == Size
}

func TypeFromId(id string) uint8 {
	switch id {
	case "i8", "sbyte":
		return I8
	case "i16":
		return I16
	case "i32":
		return I32
	case "i64":
		return I64
	case "u8", "byte":
		return U8
	case "u16":
		return U16
	case "u32":
		return U32
	case "u64":
		return U64
	case "str":
		return Str
	case "bool":
		return Bool
	case "f32":
		return F32
	case "f64":
		return F64
	case "any":
		return Any
	case "rune":
		return Rune
	case "size":
		return Size
	}
	return 0
}

func CxxTypeIdFromType(typeCode uint8) string {
	switch typeCode {
	case Void:
		return "void"
	case I8:
		return "i8"
	case I16:
		return "i16"
	case I32:
		return "i32"
	case I64:
		return "i64"
	case U8:
		return "u8"
	case U16:
		return "u16"
	case U32:
		return "u32"
	case U64:
		return "u64"
	case Bool:
		return "bool"
	case F32:
		return "f32"
	case F64:
		return "f64"
	case Any:
		return "any"
	case Str:
		return "str"
	case Rune:
		return "rune"
	case Size:
		return "size"
	}
	return ""
}

func DefaultValOfType(code uint8) string {
	if IsNumericType(code) {
		return "0"
	}
	switch code {
	case Bool:
		return "false"
	case Str:
		return `""`
	case Rune:
		return `'\0'`
	}
	return "nil"
}
