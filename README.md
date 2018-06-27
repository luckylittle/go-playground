# GO LANG

##### ...or go home :-)

---

## Introduction

1. Configure `$GOPATH`, e.g. `mkdir -p ~/Projects/go/{bin,pkg,src};mkdir -p ~/Projects/go/src/github.com/luckylittle;echo export 'GOPATH=/home/lmaly/Projects/go' >> ~/.zshrc`
2. Write Hello World in `main.go` in `$GOPATH/src/github.com/luckylittle/hello/main.go`
3. `go build` and `./hello` inside `$GOPATH/src/github.com/luckylittle/hello/`, cleanup with `go clean`
    or `go run main.go` which does not build `./hello` so you don't need to do `go clean`

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

---

[@luckylittle](https://github.com/luckylittle)
