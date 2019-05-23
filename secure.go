package secure

import "unsafe"

var dummyString = "********"
var marsharDummy = "\"" + dummyString + "\""

func SetDummyString(s string) {
	dummyString = s
	marsharDummy = "\"" + s + "\""
}

type String string

func (s String) String() string {
	return dummyString
}

func (s String) GoString() string {
	return dummyString
}

func (s String) MarshalJSON() ([]byte, error) {
	return *(*[]byte)(unsafe.Pointer(&marsharDummy)), nil
}
