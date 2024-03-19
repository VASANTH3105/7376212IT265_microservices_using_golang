package main

import "fmt"

func main() {
    text := "VASANTH"
    for i := 0; i < len(text); i++ {
        fmt.Println(string(text[i]))
    }
}
