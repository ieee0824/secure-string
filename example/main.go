package main

import (
	"fmt"

	secure "github.com/ieee0824/secure-string"
)

func main() {
	s := secure.String("password")

	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
}
