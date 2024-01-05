create a package html_generator

create a file generator.go

```go
package html_generator

import "fmt"

type Content interface {
	GetHeader() string
	Getbody() string
}

func GenerateHTML(content Content) string {

	html := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n<title>%s</title>\n</head>\n<body>\n<h1>%s</h1>\n<p>%s</p>\n</body>\n</html>", content.GetHeader(), content.GetHeader(), content.Getbody())
	return html
}


```

main.go

```go
package main

import (
	"fmt"

	"exmaple.com/typeasserrt/html_generator"
)

type myContent struct {
	heading string
	body    string
}

func (c myContent) GetHeader() string {
	return c.heading
}

func (c myContent) Getbody() string {
	return c.body
}

func main() {

	content := myContent{
		heading: "this is my heading",
		body:    "hello World!",
	}
	htmlContent := html_generator.GenerateHTML(content)

	fmt.Println(htmlContent)

}

```


Your package became very popular and now youg gets a lot of feature requests to add footer also to the html output now you changes you package and adds getFooter mehod to your interface

```go
package html_generator

import "fmt"

type Content interface {
	GetHeader() string
	Getbody() string
    Getfooter() string
}

func GenerateHTML(content Content) string {

	html := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n<title>%s</title>\n</head>\n<body>\n<h1>%s</h1>\n<p>%s</p>\n</body>\n<h2>%s</h2>\n</html>", content.GetHeader(), content.GetHeader(), content.Getbody(), content.Getfooter())
	return html
}



```

This now breaks the backward comptability of your package and break the code every where
so how will you impliment this with our breaking the backward compatibility


```go
package html_generator

import "fmt"

type Content interface {
	GetHeader() string
	GetBody() string
}

// you will create a new interface with the new method
type ContentWithFooter interface {
	GetFooter() string
	GetHeader() string
	GetBody() string
}

func GenerateHTML(content Content) string {

	if c, ok := content.(ContentWithFooter); ok {
		html := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n<title>%s</title>\n</head>\n<body>\n<h1>%s</h1>\n<p>%s</p>\n</body>\n<h2>%s</h2>\n</html>", c.GetHeader(), c.GetHeader(), c.GetBody(), c.GetFooter())
		return html
	}

	html := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n<title>%s</title>\n</head>\n<body>\n<h1>%s</h1>\n<p>%s</p>\n</body>\n</html>", content.GetHeader(), content.GetHeader(), content.GetBody())
	return html
}


```

main.go

```go
package main

import (
	"fmt"

	"exmaple.com/typeasserrt/html_generator"
)

type myContent struct {
	heading string
	body    string
	footer  string
}

func (c myContent) GetHeader() string {
	return c.heading
}

func (c myContent) GetBody() string {
	return c.body
}

//even if you comment this it will work
func (c myContent) GetFooter() string {
	return c.footer
}

func main() {

	content := myContent{
		heading: "this is my heading",
		body:    "hello World!",
		footer:  "This is my footer!",
	}
	htmlContent := html_generator.GenerateHTML(content)

	fmt.Println(htmlContent)

}


```