package main

import "fmt"

type book string

func (b book) printTitle() {
    fmt.Println(b)
}

func testing() {
    var b book = "Harry Potter"
    b.printTitle()
}