package main

import(
    "fmt
    os"
)

func main() {
    var s string
    s = os.Args[0]
    fmt.Println(s)
} 