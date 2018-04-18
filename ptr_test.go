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

func TestPBool(t *testing.T) {
	a := ptr.PBool(false)
	if *a == true {
		t.Fatal(`*a = true, wanted false`)
	}
	a = ptr.PBool(true)
	if *a == false {
		t.Fatal(`*a = false, wanted true`)
	}
}
