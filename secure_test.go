package secure

import (
	"fmt"
	"testing"
)

func TestString_String(t *testing.T) {
	s := String("foo")

	if fmt.Sprint(s) != "********" {
		t.Fatalf("failed test %v", s)
	}

	if fmt.Sprintf("%s", s) != "********" {
		t.Fatalf("failed test %s", s)
	}

	if fmt.Sprint(&s) != "********" {
		t.Fatalf("failed test %v", s)
	}

	if fmt.Sprintf("%s", &s) != "********" {
		t.Fatalf("failed test %s", s)
	}
}

func TestString_GoString(t *testing.T) {
	s := String("foo")

	if fmt.Sprint(s) != "********" {
		t.Fatalf("failed test %v", s)
	}

	if fmt.Sprintf("%v", s) != "********" {
		t.Fatalf("failed test %v", s)
	}

	if fmt.Sprint(&s) != "********" {
		t.Fatalf("failed test %v", s)
	}

	if fmt.Sprintf("%v", &s) != "********" {
		t.Fatalf("failed test %v", s)
	}
}
