package main

import (
    "fmt"
)

func showMemoryAddress(x *int) {
    fmt.Println(x)          // fmt.Println(*x) would show the value
    return
}

func main() {
    i := 1
    fmt.Println(&i)         // 0xc4200120e8
    showMemoryAddress(&i)   // Pointer to an int = 0xc4200120e8
}
