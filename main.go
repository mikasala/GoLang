package main

import (
    "fmt"
    "time"
    "github.com/mikasala/GoLang/oop"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")

    user := oop.User{}
    user.FirstName = "Maru"
    user.LastName = "Kasala"
    user.Cars = 20
    user.CarsSold = 10
    user.CarsRemaining()

}