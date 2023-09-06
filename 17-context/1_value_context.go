package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	//ctx := context.TODO()  // if you dont know what to do
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	go doSomething(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("Main done")
}
