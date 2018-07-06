# GO LANG

##### ...or go home :-)

---

## Introduction

### $GOPATH

* Defaults to your home directory like so `~/go/`.

* It’s easiest to think of the $GOPATH as being a workspace that contains all of your projects and dependencies.

* As long as all your go code lives inside it you are going to have an easier me working with go.

* You should set it yourself, so you can control the location and eventually share among multiple users:

1. Configure `$GOPATH`, e.g. `mkdir -p ~/Projects/go/{bin,pkg,src};mkdir -p ~/Projects/go/src/github.com/luckylittle;echo export 'GOPATH=/home/lmaly/Projects/go' >> ~/.zshrc`.
2. Check the Go environment variables with `go env`.

* If you look in the $GOPATH directory you will see three directories, `bin`, `pkg` and `src`. The `bin` directory is where installed binaries that have a main package are placed when you run `go install`. If you install a project without a main package it is moved into the `pkg` directory. Lastly `src` which is where you should do all your development work.

* It is good idea to `PATH=$PATH:$(go env GOPATH)/bin`

### Dependencies

* When you use the command `go get github.com/user/repo/` what happens is that a copy of the github repository is cloned into your `$GOPATH/src/github/user/repo/` directory. `go get` works with other source control systems as well so you are not limited to git.

* Inside the root of your go project dependencies are checked to exist inside the vendor directory before looking at `$GOPATH`. As such if you place your dependencies in there they will be the ones your code builds and links against. To move them into the location is thankfully simple. Install `dep` then run `dep ensure` which will inspect your code and move what is required into `vendor`. Check the `dep` docs for details on how to update/remove dependencies. Keep in mind that `dep` will place a `Gopkg.lock` and `Gopkg.toml` file in the root path for tracking these.

[Go dep](https://github.com/golang/dep)

### Building

1. Write Hello World in `main.go` in `$GOPATH/src/github.com/luckylittle/hello/main.go`.
2. `go build` and `./hello` inside `$GOPATH/src/github.com/luckylittle/hello/`, cleanup with `go clean`.
    or `go run main.go` which does not build `./hello` so you don't need to do `go clean`.

So, if you have main package:

```bash
go build   # builds the command and leaves the result in the current working directory.
go install # builds the command in a temporary directory then moves it to $GOPATH/bin.
```

For packages:

```bash
go build   # builds your package then discards the results.
go install # builds then installs the package in your $GOPATH/pkg directory.
```

If you want to cross compile, that is build on Linux for Windows or vice versa you can set what architecture your want to target and the OS through environment variables. You can view your defaults in go env but to change them you would do something like:

```bash
GOOS=darwin GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=linux GOARCH=amd64 go build
```

OS Specific source code:

```
main_darwin.go
main_linux.go
main_windows.go
```

* Multi-stage Docker build example:

```Dockerfile
FROM golang:1.10
COPY . /go/src/github.com/luckylittle/hello
WORKDIR /go/src/github.com/luckylittle/hello
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/src/github.com/luckylittle/hello .
CMD ["./main"]
```

### Testing

* To create a test file you need only create a file with `_test` as a suffix in the name. For example to create a file test you may call the file `file_test.go`.

* To run all the tests:

```bash
go test ./...
```

## Static & Dynamic Types

```go
addition = x + y

x = 1
y = "three"
```

Statically typed = !error!
Dynamically typed = "1three"

### Declaring types in Go

#### Integer

`var i int = 3`

Note: Signed int, which means positive & negative numbers are supported.

#### Float

`var f float64 = 0.111`

#### String

`var s string = "foo"`

#### Arrays

* Length and type must be specified

`var beatles [4]string`

Note: Starts with `0`.

### Check the type of variable

```go
import (
    "reflect"
)
function main () {
    ...
        fmt.Println(reflect.TypeOf(s))
        ...
}
```

### Converting between types

* For converting from/to strings, use `strconv` package from Standard Library (SL)

1. String -> Boolean:

```go
    var s string = "true"
    b, err := strconv.ParseBool(s)
```

2. Boolean -> String:

```go
    s := strconv.FormatBool(true)
    fmt.Println(s)
```

## Variables

* The type of variable is important

* After a variable is declared with a type, it is NOT possible to declare it again

* Re-assigning the value is allowed

### Shorthand declaration

`var s, t string = "foo", "bar"`
`s := "Hello World" // Only inside functions!`

### Zero values

```go
func main() {
    var i int       // i = 0
    var f float64   // f = 0
    var b bool      // b = false
    var s string    // s = ""
}
```

* Checking if the zero value has been assigned?

```go
    func main() {
        var s string    // s = ""
        if s == "" {
            fmt.Printf("s has not been assigned a value and is zero valued")
        }
    }
```

### Variable scope

* Lexically scoped using blocks = meaning Go defines where variables can or cannot be referenced.

* A block is defined as a possible empty sequence of declarations and statements within matching brace brackets = `{` & `}`.

* A variable declared within those brackets may be accessed anywhere within this block.

* Brackets within brackets denote a new, `inner` block.

* `Inner` blocks can access vars within `outer` blocks.

* `Outer` blocks CANNOT access vars within `inner` blocks.

### Pointers

* Allocated position in computer's memory = `&variable`

* Prevent two instances of the variable in different memory locations

* By using asterisk, the value is printed = `*variable`

```go
func main() {
    s := "Hello world"
    fmt.Println(&s) // 0xc42000e1e0
}
```

### Constants

* Values do not change during the life of a program

`const greeting string = "Hello, world"`

## Functions

* Function signature:

```go
func addUp(x int, y int) int { // expected variables in the brackets (), after closing bracket comes the return value
    return x + y               // if the func signature declares a return value, the func body must end in a terminating statement
}                              // terminated function body
```

### Returning multiple values

```go
func getPrize() (int, string) {
    i := 2
    s := "goldfish:
    return i, s
}

func main () {
    quantity, prize := getPrize()
    fmt.Printf("You won %v %v\n", quantity, prize)
}
```

### Variadic func

* Accept variable number of arguments using 3 dots (`...`)

```go
func sumNumbers(numbers...int) int {
    ...
}
```

### Named return values

* Assign values to named variables befre they are returned

* The func signature declares the variables as part of the return values

```go
func sayHi() (x, y string) {
    x = "hello"
    y = "world"
    return                      // naked return statement, works if you use named return values. Returns the named variables in the same order.
}
```

### Recursive functions

* Can call themselves indefinitely or until particular condition is met

* Calls itself as the result value of a terminating statement

### Passing functions as values

* It is possible to assign functions to a value and call them at a later date

* Functions are type in Go so they can be passed to another function

```go
func main () {
    fn := func() {
        fmt.Println("function called")
    }
    fn() // function is called
}
```

Passing func as an argument:

```go
/* recursive function */
package main

import "fmt"

func anotherFunction(f func() string) string {  // sub function signature shows funct argument that returns string, receiving function also returns string
    return f()
}

func main() {
    fn := func() string {
        return "function called"
    }
    fmt.Println(anotherFunction(fn))
}
```

## Control flow

### if

* Multiple `if` statements can be run one after another and they will be evaluated in the order they are in the source code.

```go
func main () {
    b := true
    if b {                              // evaluates whether b is true
        fmt.Println("b is true!")       // code is executed, because b is true
    }
}
```

### else

* If nothing else is true, run this.

```go
func main () {
    b := false
    if b {
        fmt.Println("b is true!")
    } else {
        fmt.Println("b is false!")
    }
}
```

### else if

```go
func main() {
    i := 2
    if i == 3 {
        fmt.Println("i is 3")
    } else if i == 2 {
        fmt.Println("i is 2")
    }
}
```

### Comparison, arithmetic operators, logical operators

* Must be the same type

* `==`, `!=`, `<`, `<=`, `>`, `>=`

* `+`, `-`, `*`, `/`, `%`

* `&&`, `||`, `!`

### switch

```go
func main() {
    i := 2
    switch i {
        case 2:                     // if an expression is found to be true, code is evaluated
            fmt.Println("Two")
        case 3:
            fmt.Println("Three")
        case 4:
            fmt.Println("Four")
        default:                    // if none of the case statements is true
            fmt.Println("I don't know")
    }
}
```

### for

```go
func main() {
    i := 0
    for i < 10 {                  // if this is true, the code is executed
        i++                       // incremented by one
        fmt.Println("i is", i)
    }                             // when i is no longer less than 10, no longer executed and loop stops
}
```

## For with init & post

* `init` = this is run only once before the first iteration

* `post` = this is evaluated after each iteration

```go
func main() {
    for i := 0;i <10;i++ {
        fmt.Println("i is", i)
    }
}
```

### For with range

* Can loop over data structure

* Iteration starts at `0`

```go
func main() {
    numbers := []int{1,2,3,4}                        // slice containing four integers
    for i, n := range numbers {                      // iteration variable = i, value = n
        fmt.Println("The index of the loop is", i)
        fmt.Println("The value from the array is", n)
    }
}
```

### defer

* Allows a func to be executed after surrounding func returns.

* Usually cleanup operations.

* If multiple `defer`s are present, they will be executed in reverse order that they were declared in the source code.

```go
func main(){
    defer fmt.Println("I am run after the function completes")      // last
    fmt.Println("Hello World!")                                     // first
}
```

---

[@luckylittle](https://github.com/luckylittle)
