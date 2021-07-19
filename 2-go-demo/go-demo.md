1. create foldter go-training
2. create go-training./go-demo
3. go to terminal and type go mod init cybage.com/go-demo
4. create folder mascot
5. create file mascot/mascost.go
6. This is going to trigger up the go extension .
7. Click on analysis tools and click install
8. in mascot.go

```
package mascot

func BestMascot() string {
	return "Tux"
}
```

9. main.go 
    
```
package main

import (
	"fmt"

	"cybage.com/go-demo/mascot"
)

func main() {
	fmt.Println(mascot.BestMascot())
}
```

10. let us see how to use third party package 

11.  in main.go add import the rsc.io/quote package and add a call to its Go function
```
fmt.Println(quote.Go())
```
12.  go mod tidy
13.  how to write a test
Go has a built-in testing command called go test and a package testing which combine to give a minimal but complete testing experience.
14. create file mascot_test.go
```
package mascot

import "testing"

func TestBestMascot(t *testing.T) {
	if BestMascot() != "Gopher" {
		t.Fatal("wrong mascot!")
	}
}
```
15.  cd mascot
16.  go test 
17.  go test tool has built-in code-coverage go test -cover


https://golang.org/doc/tutorial/getting-started