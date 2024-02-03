package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    usernameFile, err := os.Open("/etc/secrets/username")
    if err != nil {
        panic(err)
    }
    defer usernameFile.Close()

    passwordFile, err := os.Open("/etc/secrets/password")
    if err != nil {
        panic(err)
    }
    defer passwordFile.Close()

    usernameContent, err := io.ReadAll(usernameFile)
    if err != nil {
        panic(err)
    }

    passwordContent, err := io.ReadAll(passwordFile)
    if err != nil {
        panic(err)
    }

    fmt.Println("Username:", string(usernameContent))
    fmt.Println("Password:", string(passwordContent))
}
