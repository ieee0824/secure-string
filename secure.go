package secure

var dummyString = "********"

func SetDummyString(s string) {
	dummyString = s
}

type String string

func (s String) String() string {
	return dummyString
}

func (s String) GoString() string {
	return dummyString
}
