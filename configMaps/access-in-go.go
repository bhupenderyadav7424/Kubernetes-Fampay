package main

import (
    "os"
    "fmt"
)

func main() {
    key1 := os.Getenv("key1")
    key2 := os.Getenv("key2")
    key3 := os.Getenv("key3")

    fmt.Println("key1:", key1)
    fmt.Println("key2:", key2)
    fmt.Println("key3:", key3)
}
