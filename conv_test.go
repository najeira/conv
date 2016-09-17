package conv

import (
	"database/sql"
	"testing"
)

func TestStringString(t *testing.T) {
	var s string = "hoge"
	if r := String(interface{}(s)); r != s {
		t.Errorf("String invalid %v", r)
	}
}

func TestStringStringPointer(t *testing.T) {
	var s string = "hoge"
	if r := String(interface{}(&s)); r != s {
		t.Errorf("String invalid %v", r)
	}
}

func TestStringNullString(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "hoge"}
	if r := String(interface{}(s)); r != s.String {
		t.Errorf("String invalid %v", r)
	}
}

func TestStringNullStringPointer(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "hoge"}
	if r := String(interface{}(&s)); r != s.String {
		t.Errorf("String invalid %v", r)
	}
}

func TestStringBytes(t *testing.T) {
	var s []byte = []byte("hoge")
	if r := String(interface{}(s)); r != "hoge" {
		t.Errorf("String invalid %v", r)
	}
}

func TestStringRune(t *testing.T) {
	var s = 'r'
	if r := String(interface{}(s)); r != "r" {
		t.Errorf("String invalid %v", r)
	}
}

func TestStringRuneSlice(t *testing.T) {
	var s = []rune{'1', '2', '3'}
	if r := String(interface{}(s)); r != "123" {
		t.Errorf("String invalid %v", r)
	}
}

func TestIntInt64(t *testing.T) {
	var s int64 = 123
	if r := Int(interface{}(s)); r != s {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt32(t *testing.T) {
	var s int32 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt16(t *testing.T) {
	var s int16 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt8(t *testing.T) {
	var s int8 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt(t *testing.T) {
	var s int = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt64Pointer(t *testing.T) {
	var s int64 = 123
	if r := Int(interface{}(&s)); r != s {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt32Pointer(t *testing.T) {
	var s int32 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt16Pointer(t *testing.T) {
	var s int16 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntInt8Pointer(t *testing.T) {
	var s int8 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntIntPointer(t *testing.T) {
	var s int = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint64(t *testing.T) {
	var s uint64 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint32(t *testing.T) {
	var s uint32 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint16(t *testing.T) {
	var s uint16 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint8(t *testing.T) {
	var s uint8 = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint(t *testing.T) {
	var s uint = 123
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint64Pointer(t *testing.T) {
	var s uint64 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint32Pointer(t *testing.T) {
	var s uint32 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint16Pointer(t *testing.T) {
	var s uint16 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUint8Pointer(t *testing.T) {
	var s uint8 = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntUintPointer(t *testing.T) {
	var s uint = 123
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntFloat64(t *testing.T) {
	var s float64 = 123.456
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntFloat32(t *testing.T) {
	var s float32 = 123.456
	if r := Int(interface{}(s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntFloat64Pointer(t *testing.T) {
	var s float64 = 123.456
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntFloat32Pointer(t *testing.T) {
	var s float32 = 123.456
	if r := Int(interface{}(&s)); r != int64(s) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntNullInt64(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if r := Int(interface{}(s)); r != s.Int64 {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntNullFloat64(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if r := Int(interface{}(s)); r != int64(s.Float64) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntNullString(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if r := Int(interface{}(s)); r != int64(123) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntNullInt64Pointer(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if r := Int(interface{}(&s)); r != s.Int64 {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntNullFloat64Pointer(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if r := Int(interface{}(&s)); r != int64(s.Float64) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestIntNullStringPointer(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if r := Int(interface{}(&s)); r != int64(123) {
		t.Errorf("Int invalid %v", r)
	}
}

func TestFloatInt64(t *testing.T) {
	var s int64 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt32(t *testing.T) {
	var s int32 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt16(t *testing.T) {
	var s int16 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt8(t *testing.T) {
	var s int8 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt(t *testing.T) {
	var s int = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint64(t *testing.T) {
	var s uint64 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint32(t *testing.T) {
	var s uint32 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint16(t *testing.T) {
	var s uint16 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint8(t *testing.T) {
	var s uint8 = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint(t *testing.T) {
	var s uint = 123
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatNullInt64(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if r := Float(interface{}(s)); r != float64(s.Int64) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatNullFloat64(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if r := Float(interface{}(s)); r != s.Float64 {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatNullString(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if r := Float(interface{}(s)); r != float64(123) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt64Pointer(t *testing.T) {
	var s int64 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt32Pointer(t *testing.T) {
	var s int32 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt16Pointer(t *testing.T) {
	var s int16 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatInt8Pointer(t *testing.T) {
	var s int8 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatIntPointer(t *testing.T) {
	var s int = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint64Pointer(t *testing.T) {
	var s uint64 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint32Pointer(t *testing.T) {
	var s uint32 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint16Pointer(t *testing.T) {
	var s uint16 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUint8Pointer(t *testing.T) {
	var s uint8 = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatUintPointer(t *testing.T) {
	var s uint = 123
	if r := Float(interface{}(&s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatFloat64(t *testing.T) {
	var s float64 = 123.456
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatFloat32(t *testing.T) {
	var s float32 = 123.456
	if r := Float(interface{}(s)); r != float64(s) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatNullInt64Pointer(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if r := Float(interface{}(&s)); r != float64(s.Int64) {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatNullFloat64Pointer(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if r := Float(interface{}(&s)); r != s.Float64 {
		t.Errorf("Float invalid %v", r)
	}
}

func TestFloatNullStringPointer(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if r := Float(interface{}(&s)); r != float64(123) {
		t.Errorf("Float invalid %v", r)
	}
}
