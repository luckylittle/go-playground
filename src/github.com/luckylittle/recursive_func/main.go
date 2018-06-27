/* recursive function */
package main

import "fmt"

func feedMe(portion int, eaten int) int {
    eaten = portion + eaten
    if eaten >= 5 {
        fmt.Printf("I'm full! I've eaten %d\n", eaten)
        return eaten
    }
    fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
    return feedMe(portion, eaten)   // function calls itself and the function will be executed again
}

func main() {
    feedMe(1, 0)
}
