package main

import (
    "fmt"
)

var s = "Hello world" // Anything outside of the first level of brace brackets in a file is available within all blocks

func main() {
    fmt.Printf("Printing `s` variable from outer block %v\n", s)
    b := true
    if b {
        fmt.Printf("Printing `b` variable from outer block %v\n", b)
        i := 42
        if b != false {
            fmt.Printf("Printing `i` variable from outer block %v\n", i)
        }
    }
    // Printing `s` variable from outer block Hello world
    // Printing `b` variable from outer block true
    // Printing `i` variable from outer block 42
}
