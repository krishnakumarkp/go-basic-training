package main

import "fmt"


// Basic types

// Built-in basic types in Go
// String, bool, numeric(int8, uint8 (byte), int16, uint16, int32 (rune), uint32, 
// int64, uint64, int, uint, uintptr, float32, float64,complex64, complex128.)

//Composite Types

// Assume T is an arbitrary type and Tkey is a type of the key

*T         // a pointer type
[5]T       // an array type
[]T        // a slice type
map[Tkey]T // a map type

// a struct type
struct {
	name string
	age  int
}

// a function type functions are first-class types in Go.
func(int) (bool, string)

// an interface type
interface {
	Method0(string) int
	Method1() (int, bool)
}

// some channel types
chan T
chan<- T
<-chan T

//Type Definitions
// in Go, we can define new types by using the following form. In the syntax, type is a keyword.

type NewTypeName SourceType

type myInt int
type mySlice []int
var number myInt 
var slice mySlice

// The following new defined and source types are all basic types.
type MyInt int
type (
	Age   int
	Text  string
)

var number MyInt
number = 10
fmt.Println(number)

// The following new defined and source types are
// all composite types.
type IntPtr *int
type Book struct{author, title string; pages int}
type Convert func( int,  bool)( int,  string)
type StringArray [5]string
type StringSlice []string

//Type Alias Declarations

// we can declare custom type aliases by using the following syntax. The syntax of alias declaration 
// is much like type definition, but please note there is a = in each type alias specification.

type Name string
type Age  = int

type table = map[string]int
type Table = map[Name]Age

func main() {
	var a Name
	var b string
	b = "test"
	a = b
}

func f() {
	// The names of the three defined types
	// can be only used within the function.
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{Read([]byte) int}
}

main() {
	var a PersonAge
}

