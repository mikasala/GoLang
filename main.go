package main

import (
    "fmt"
    "time"
    "sync"
    "net/http"
    "github.com/mikasala/GoLang/models"
    "github.com/mikasala/GoLang/api"
)

func f(word string) {
    for i := 0; i < len(word); i++ {
        fmt.Println(word, ":", i)
    }
}

func main() {

    // Go Routines

    var (
		wg sync.WaitGroup
    )

    testMsg := "Sample String"
    testMsg2 := "Sample String 2"
    wg.Add(1)
    go func(wg *sync.WaitGroup, msg *string) {// go routine
        defer wg.Done() // flag that says it there wait group process is not done yet
        
        *msg += " - 1" // append string
    }(&wg, &testMsg)// &wg : pointer to waitgroup, &testMsg : the msg string

    wg.Add(1)
    go func(wg *sync.WaitGroup, msg *string) {// go routine
        defer wg.Done()
        
        *msg += " - 2" // append string
    }(&wg, &testMsg2)

    wg.Wait()// wait for all the go routines to finish under wait group wg
    
    fmt.Println(testMsg + " : " + testMsg2)

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("end")

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
    anotherUser.SetFirstName("Another First Name") // setting the value of private property using public setter
    anotherUser.SetLastName("Another Last Name") // setting the value of private property using public setter
    anotherUser.SetCars(100) // setting the value of private property using public setter
    anotherUser.SetCarsSold(22) // setting the value of private property using public setter
    anotherUser.PrintCarsRemaining()

    user2 := models.NewUser()// new user: calls constructor-like method/function
    // user2.firstName = "you" // error since firstName is a private property (starts with small letter)
    user2.SetFirstName("you") //  set firstname to you
    user2.PrintCarsRemaining()
    user2.AppendLastNameToFirstName()
    fmt.Println(user2.GetFirstName())

    user3 := models.NewPrivateUser()// new user: calls constructor-like method/function
    user3.SetFirstName("you") //  set firstname to you
    fmt.Println(user3.GetFirstName())

}

func regHandlers(){
    http.HandleFunc("/test", api.TestAPI)
}