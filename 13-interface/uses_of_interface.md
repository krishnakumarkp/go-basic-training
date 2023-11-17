# in object oriented programming a interface is a common means for unrealted objects to communicate with each other.

```go

package main

import "fmt"

// Define the ElectronicDevice interface
type ElectronicDevice interface {
    TurnOn()
    TurnOff()
}

// Implement the interface in unrelated types
type Smartphone struct{}

func (s Smartphone) TurnOn() {
    fmt.Println("Smartphone is turning on.")
}

func (s Smartphone) TurnOff() {
    fmt.Println("Smartphone is turning off.")
}

type Laptop struct{}

func (l Laptop) TurnOn() {
    fmt.Println("Laptop is turning on.")
}

func (l Laptop) TurnOff() {
    fmt.Println("Laptop is turning off.")
}

// Define the Person struct
type Person struct{}

// Behavior for a person operating an electronic device
func (p Person) OperateElectronicDevice(device ElectronicDevice) {
    fmt.Println("Operating the electronic device:")
    device.TurnOn()
    device.TurnOff()
}

func main() {
    myPhone := Smartphone{}
    myLaptop := Laptop{}
    person := Person{}

    // person and smart phone are unrealted object ,turning on smar phone may be different from turning on laptop.
    // to make unrealted objects communicate we can define an interface.

    // Operate the smartphone
    person.OperateElectronicDevice(myPhone)

    // Operate the laptop
    person.OperateElectronicDevice(myLaptop)
}

```
# writing genric algorithms

```go
package main

import (
	"fmt"
)

// Comparer interface defines a method for comparing two values.
type Comparer interface {
	Compare(other interface{}) int
}

// Max finds the maximum value in a slice of comparables.
func Max(slice []Comparer) Comparer {
	if len(slice) == 0 {
		return nil
	}

	maxValue := slice[0]

	for _, value := range slice[1:] {
		if value.Compare(maxValue) > 0 {
			maxValue = value
		}
	}

	return maxValue
}

// Integer type implements the Comparer interface for integers.
type Integer int

// Compare compares two integers.
func (i Integer) Compare(other interface{}) int {
	otherInt := other.(Integer)
	if i > otherInt {
		return 1
	} else if i < otherInt {
		return -1
	}
	return 0
}

type Word string

// Compare compares length two words.
func (w Word) Compare(other interface{}) int {
	otherWord := other.(Word)
	wordLength := len(w)
	otherLength := len(otherWord)
	if wordLength > otherLength {
		return 1
	} else if wordLength < otherLength {
		return -1
	}
	return 0
}

func main() {
    //intiger has different algorithm for compare and words has different algorithm
	intSlice := []Comparer{
		Integer(5),
		Integer(3),
		Integer(7),
		Integer(1),
	}

	maxValue := Max(intSlice).(Integer)
	fmt.Println("Maximum value:", maxValue)

	wordSlice := []Word{
		"apple",
		"banana",
		"cherry",
		"date",
		"grape",
		"kiwi",
	}

	//https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces/12754757#12754757
	var comparerSlice []Comparer

	for _, word := range wordSlice {
		comparerSlice = append(comparerSlice, word)
	}

	longWord := Max(comparerSlice).(Word)
	fmt.Println("Maximum value:", longWord)
}


```

# hiding implementaion detials

 interfaces are often used to hide implementation details from the users of a package or module. This concept is commonly known as encapsulation or abstraction. The idea is to expose only the necessary functionality through the interface while keeping the internal details hidden. 

 In the above example, type Word implements the Comparer interface, and the Max function takes a slice of Comparer interface as a parameter. Users of the Max function only need to know about the Comparer interface and not the details of the Word implementation. This separation allows for flexibility in changing or extending the implementation of Word without affecting code that relies on the Comparer interface.


# providing interception points

call dispath

 interfaces allow a user to intercept and customize the behavior of a type by satisfying that interface. 

 Interfaces define a set of methods that a type must implement to satisfy the interface. These methods serve as interception points, allowing users to intercept or customize the behavior of a type by providing their own implementations for those methods.

 eg:

 let us assume we are using kklog.AwesomeLog  for loggin in our code

 create a package kklog/kklog.go

 ```go
 package kklog

import "fmt"

// Logger is an interface with a WriteLog method.
type Logger interface {
	WriteLog(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) WriteLog(message string) {
	fmt.Printf("Logging to console: %s\n", message)
}

type AwesomeLog struct {
	LogDriver Logger
}

func (a AwesomeLog) Log(message string) {
	a.LogDriver.WriteLog(message)
}

func NewAwesomeLog() AwesomeLog {
	default_logger := ConsoleLogger{}
	return AwesomeLog{
		LogDriver: default_logger,
	}
}
 ```

 now in main.go

 ```go
 
 func main() {

	awesomeLog := kklog.NewAwesomeLog()

	awesomeLog.Log("some error occured!")
 }
 
 ```

 this will give out put as  Logging to console: some error occured!

 what if we need to change the bhavior of kklog.AwesomeLog so that it will print the log as json?
 what if the kklog is a third party package?


 in main.go

 ```go
 //create a struct for json format
type LogFormat struct {
	Message string
}

 //create a new struct 
 type JsonDriver struct {
	defaultDriver kklog.Logger //interfae
}

//attach WriteLog method to satisfy kklog.Logger interface
func (j JsonDriver) WriteLog(message string) {
	myLogFormat := LogFormat{
		Message: message,
	}

	jsonData, err := json.Marshal(myLogFormat)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	j.defaultDriver.WriteLog(string(jsonData))

}

func main() {

	awesomeLog := kklog.NewAwesomeLog()

	awesomeLog.Log("some error occured!")

	myNewJsonDriver := JsonDriver{
		defaultDriver: awesomeLogger.LogDriver,
	}

	awesomeLogger.LogDriver = myNewJsonDriver

	awesomeLogger.Log("some other error occured!")

	//awesomeLogger->ConsoleLogger  was intercepted by our custom myNewJsonDriver
	//awesomeLogger->myNewJsonDriver->ConsoleLogger

	//if can intercept with one logger we can do with many
	//awesomeLogger->base64EncoderDriver->myNewJsonDriverr->ConsoleLogger
	//this is called interface chaining

}

 ```


# creating custom errors

create folder fileprocessor and create file fileprocessor/fileprocessor.go

```go
package fileprocessor

import (
	"errors"
	"os"
)

// ProcessFile simulates a file processing operation.
func ProcessFile(fileName string) error {
	// Simulate a file not found error
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return errors.New("File not found: " + fileName)
	}

	// Simulate a general file processing error
	return errors.New("Error processing file " + fileName + " during read operation: something went wrong")
}


```

now in main.go

```go

func main() {
	fileName := "example.txt"

	err := fileprocessor.ProcessFile(fileName)

	// Handle errors
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File processed successfully.")
	}
	//what if we need to do identify the error in above code?
	// we will need to write some bad code like this
	// this is bad becuse the error messages are for humans to read
	// not to be checked from the code

	if err != nil {

		if strings.Contains(err.Error(), "File not found") {
			fmt.Println("Was not able to locate the file that you specified.")
		}
		if strings.Contains(err.Error(), "Error processing file") {
			fmt.Println("Was not able to process the file.")
		}

	} else {
		fmt.Println("File processed successfully.")
	}

}

```
How can we code this better ?

error is a built-in interface type in Go.

```go
type error interface {
    Error() string
}
```

so any type that impliments this interface can be returned as an error


create folder fileprocessor and create file fileprocessor/fileprocessor.go

in fileprocessor.go

```go
package fileprocessor

import (
	"fmt"
	"os"
)

// FileNotFoundError is a custom error type for file not found errors.
type FileNotFoundError struct {
	FileName string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found: %s", e.FileName)
}

// FileProcessingError is a custom error type for general file processing errors.
type FileProcessingError struct {
	FileName  string
	Operation string
	Err       error
}

func (e *FileProcessingError) Error() string {
	return fmt.Sprintf("Error processing file %s during %s operation: %s", e.FileName, e.Operation, e.Err.Error())
}

// ProcessFile simulates a file processing operation.
func ProcessFile(fileName string) error {
	// Simulate a file not found error
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return &FileNotFoundError{FileName: fileName}
	}

	// Simulate a general file processing error
	return &FileProcessingError{
		FileName:  fileName,
		Operation: "read",
		Err:       fmt.Errorf("something went wrong"),
	}
}


```
in main.go
```go
package main

import (
	"fmt"
	"go-training/struct/fileprocessor"
)

func main() {
	fileName := "example.txt"

	err := fileprocessor.ProcessFile(fileName)

	// Handle custom errors
	if err != nil {
		// Check if err is not nil before entering the type switch
		switch e := err.(type) {
		case *fileprocessor.FileNotFoundError:
			fmt.Println("Was not able to locate the file that you specified:", e.FileName)
		case *fileprocessor.FileProcessingError:
			fmt.Printf("Error processing file %s during %s operation: %s\n", e.FileName, e.Operation, e.Err.Error())
		default:
			fmt.Println("Generic error:", err.Error())
		}
	} else {
		fmt.Println("No error. File processed successfully.")
	}
}

```