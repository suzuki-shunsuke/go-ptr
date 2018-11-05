package ptr_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"
)

func TestPStr(t *testing.T) {
	a := ptr.PStr("hello")
	if *a != "hello" {
		t.Fatalf(`*a = "%s", wanted "hello"`, *a)
	}
}

func TestPInt(t *testing.T) {
	a := ptr.PInt(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPUint(t *testing.T) {
	a := ptr.PUint(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPBool(t *testing.T) {
	a := ptr.PBool(false)
	if *a {
		t.Fatal(`*a = true, wanted false`)
	}
	a = ptr.PBool(true)
	if !*a {
		t.Fatal(`*a = false, wanted true`)
	}
}

func TestPUint8(t *testing.T) {
	a := ptr.PUint8(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPUint16(t *testing.T) {
	a := ptr.PUint16(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPUint32(t *testing.T) {
	a := ptr.PUint32(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPUint64(t *testing.T) {
	a := ptr.PUint64(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPInt8(t *testing.T) {
	a := ptr.PInt8(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPInt16(t *testing.T) {
	a := ptr.PInt16(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPInt32(t *testing.T) {
	a := ptr.PInt32(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPInt64(t *testing.T) {
	a := ptr.PInt64(5)
	if *a != 5 {
		t.Fatalf(`*a = %d, wanted 5`, *a)
	}
}

func TestPFloat32(t *testing.T) {
	a := ptr.PFloat32(5)
	if *a != 5 {
		t.Fatalf(`*a = %f, wanted 5`, *a)
	}
}

func TestPFloat64(t *testing.T) {
	a := ptr.PFloat64(5)
	if *a != 5 {
		t.Fatalf(`*a = %f, wanted 5`, *a)
	}
}

func TestPComplex64(t *testing.T) {
	a := ptr.PComplex64(5)
	if *a != 5 {
		t.Fatalf(`*a = %f, wanted 5`, *a)
	}
}

func TestPComplex128(t *testing.T) {
	a := ptr.PComplex128(5)
	if *a != 5 {
		t.Fatalf(`*a = %f, wanted 5`, *a)
	}
}

func TestPRune(t *testing.T) {
	a := ptr.PRune('a')
	if *a != 'a' {
		t.Fatalf(`*a = %d, wanted 'a'`, *a)
	}
}

func TestPByte(t *testing.T) {
	a := ptr.PByte(byte('a'))
	if *a != byte('a') {
		t.Fatalf(`*a = %d, wanted byte('a')`, *a)
	}
}
