# secure-string

Output confidential string to stdout as dummy string.

## example

```Go
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
```

```
$ go run example/main.go 
********
********
```