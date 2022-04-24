package mytypes_test

import (
	"bookstore/mytypes"
	"testing"
)

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder

	mb.WriteString("Hello, ")
	mb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.String()

	if want != got {
		t.Errorf("want %q, got %q", want , got)
	}
	wantLen := 15
	gotLen := mb.Len()

	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d", mb.String(), wantLen, gotLen)
	}
}

func TestToUppercase(t *testing.T){
	t.Parallel()
	var su mytypes.StringUppercaser
	su.WriteString("Hi ")
	su.WriteString("World")
	want := "Hi World"
	got := su.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

	wantUpperCase := "HI WORLD"
	gotUppercase := su.ToUpper()

	if wantUpperCase != gotUppercase {
		t.Errorf("ToUpper(%q): want %q, got %q", su.String(), wantUpperCase, gotUppercase)
	}
}


func TestDouble(t *testing.T){
	t.Parallel()
	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)
	x.Double()
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}

