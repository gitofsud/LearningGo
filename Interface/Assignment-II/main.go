package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    filename := strings.Join(os.Args[1:], "")
	file, err := os.Open(filename)
	
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    fmt.Println("Reading file:", filename)
    io.Copy(os.Stdout, file)
}