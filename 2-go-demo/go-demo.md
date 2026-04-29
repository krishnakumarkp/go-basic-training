
# 🟢 Go Modules Hands-on 

## 1. Create Project Structure

```bash
mkdir go-training
cd go-training
mkdir go-demo
cd go-demo
```

---

## 2. Initialize Go Module

```bash
go mod init github.com/krishnakumarkp/go-training/go-demo
```

👉 This creates a `go.mod` file (no need for GOPATH setup).

---

## 3. Create Package

```bash
mkdir mascot
touch mascot/mascot.go
```

---

## 4. Write Code in `mascot/mascot.go`

```go
package mascot

func BestMascot() string {
	return "Gopher"
}
```

---

## 5. Create `main.go`

```go
package main

import (
	"fmt"

	"github.com/krishnakumarkp/go-training/go-demo/mascot"
)

func main() {
	fmt.Println(mascot.BestMascot())
}
```

---

## 6. Run the Program

```bash
go run main.go
```

---

## 7. Add a Third-Party Package

Update `main.go`:

```go
package main

import (
	"fmt"

	"github.com/krishnakumarkp/go-training/go-demo/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}
```

---

## 8. Download Dependencies

```bash
go mod tidy
```

👉 This:

* Downloads dependencies
* Updates `go.mod` and `go.sum`

---

## 9. Write Unit Test

Create file:

```bash
touch mascot/mascot_test.go
```

### `mascot_test.go`

```go
package mascot

import "testing"

func TestBestMascot(t *testing.T) {
	if BestMascot() != "Gopher" {
		t.Fatal("wrong mascot!")
	}
}
```

---

## 10. Run Tests

```bash
cd mascot
go test
```

---

## 11. Run Tests with Coverage

```bash
go test -cover
```

---

## 🔗 Reference

[https://go.dev/doc/tutorial/getting-started](https://go.dev/doc/tutorial/getting-started)

