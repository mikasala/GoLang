package main

import (
    "fmt"
    "time"
    "github.com/mikasala/GoLang/models"
)

func f(word string) {
    for i := 0; i < len(word); i++ {
        fmt.Println(word, ":", i)
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

    // user := models.User{}
    // user.FirstName = "Maru"
    // user.LastName = "Kasala"
    // user.Cars = 20
    // user.CarsSold = 10
    // user.PrintCarsRemaining()

    // public model with public properties
    publicUser := models.PublicUser{}
    publicUser.FirstName = "you" // setting value of public property
    fmt.Println(publicUser.FirstName)

    // public model with private properties
    anotherUser := models.User{}
    anotherUser.SetFirstName("you") // setting the value of private property using public public setter
    anotherUser.PrintCarsRemaining()

    user2 := models.NewUser()// new user: calls constructor-like method/function
    // user2.firstName = "you" // error since firstName is a private property (starts with small letter)
    user2.SetFirstName("you") //  set firstname to you
    user2.PrintCarsRemaining()
    user2.AppendLastNameToFirstName()
    fmt.Println(user2.GetFirstName())

}