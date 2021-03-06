package main

import (
	"encoding/json"
	"fmt"
	"os"

	secure "github.com/ieee0824/secure-string"
)

type userInfo struct {
	ID       int
	Password secure.String
}

func main() {
	user := &userInfo{
		1,
		"foo",
	}

	fmt.Printf("%v\n", user)

	json.NewEncoder(os.Stdout).Encode(user)
}
