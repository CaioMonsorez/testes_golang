// how to create a calculator on golang?    
package main

import (
    "fmt"
)

func add(m ...int) int {
    sum := 0
    for _, i := range m {
        sum += i
    }
    return sum
}
func sub(m ...int) int {
    sub := m[0]
    for _, i := range m[1:] {
        sub -= i
    }
    return sub
}
func mul(m ...float32) float32 {
    c := float32(1)
    for _, i := range m {
        c *= i
    }
    return c
}
func div(m ...float32) float32 {
    quo := m[0]
    for _, i := range m[1:] {
        quo /= i
    }
    return quo
}

func main() {
    fmt.Println(add(4, 6))
    fmt.Println(sub(4, 6))
    fmt.Println(mul(4, 6))
    fmt.Println(div(6, 4))
}



