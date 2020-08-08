package oop

import (
	"fmt"
)

type User struct {
    FirstName 		string
	LastName 		string
	Cars 			int
	CarsSold		int
}

func (u User) CarsRemaining() {
	fmt.Printf("%s %s has %d cars\n", u.FirstName, u.LastName, (u.Cars - u.CarsSold))
}

