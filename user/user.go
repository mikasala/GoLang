// In C# : namespace user{class user}
package user

// In C#: using "fmt"
import (
	"fmt"
)

// In Java/PHP: uses this.[private variable] to access private variables
// In C# : just directly use private variable

// class User{
//  // In PHP:
//	// private $__firstName
// 	private string __firstName;
// 	private string __lastName;
// 	private int __cars;
// 	private int __carsSold;

//	// Constructor
// 	// In PHP:
// 	//function __construct() {
// 	//	$this->__firstName 	= "First Name";
// 	//	$this->__lastName 	= "Last Name";
// 	//	$this->__cars 		= 0;
// 	//	$this->__carsSold	= 0;
// 	//}
// 	public User(){
// 		__firstName = "First Name";
// 		__lastName 	= "Last Name";
// 		__cars 		= 0;
// 		__carsSold	= 0;
// 	}

// 	// Properties

// 	// In Java:
// 	// public String getFirstName(){
// 	// 	return this.__firstName;
// 	// }
// 	// public void setFirstName(String value){
// 	// 	this.__firstName = value;
// 	// }

// 	// In PHP:
// 	// public function getFirstName(){
// 	// 	return $this->__firstName;
// 	// }
// 	// public function setFirstName(String value){
// 	// 	$this->__firstName = value;
// 	// }

// 	public string FirstName {
// 		get { return __firstName; }
// 		set { __firstName = value; }
// 	}

// 	public string LastName {
// 		get { return __lastName; }
// 		set { __lastName = value; }
// 	}

// 	public int Cars {
// 		get { return __cars; }
// 		set { __cars = value; }
// 	}

// 	public int CarsSold {
// 		get { return __carsSold; }
// 		set { __carsSold = value; }
// 	}

// 	// Methods 
//	// In PHP
//	// public function PrintCarsRemaining(){
//	//	echo $this->__firstName + " " + $this->__lastName + " has " + ($this->__cars - $this->__carsSold) + "cars\n";
//	// }
// 	public void PrintCarsRemaining(){
// 		Console.WriteLine(__firstName + " " + __lastName + " has " + (__cars - __carsSold) + "cars\n");
// 	}
// }

// In C/C#/Java: 		[data type] [variable name]
// In Golang: 	[variable name] [data type]

// Class/Object
// Case Sensitive
// Uppercase : exported class/struct
// Can be different from file name
type User struct {
    FirstName 		string
	LastName 		string
	Cars 			int
	CarsSold		int
}

// Lowercase : imported class/struct
// if using imported class/struct,
// the struct name must be same as folder name,
// unlike exported, any exported struct name can be used
type user struct {
    firstName 		string
	lastName 		string
	cars 			int
	carsSold		int
}

// Functions with return data:
// In C/C#/Java: 	[data type] [functon name]([parameters])
// In Golang: 		[function name]([parameters]) [data type]

// Functions of void type:
// In C/C#/Java: 	void [functon name]([parameters])
// In Golang: 		[function name]([parameters])

// New() is the function ; user is the type of data to be returned
func New() user{// Acts like constructor
	user := user{}
	user.firstName = "First Name"
	user.lastName = "Last Name"
	user.cars = 0
	user.carsSold = 0

	return user
}

// In C/Golang: Needs to declare the ([variable] [data type]) that calls the function with
// to be able to call it as its(the data type) method

// Pass by Value
// This works like methods with type void
func (u User) PrintCarsRemaining() {
	fmt.Printf("%s %s has %d cars\n", u.FirstName, u.LastName, (u.Cars - u.CarsSold))
}
func (u user) PrintCarsRemaining() {
	fmt.Printf("%s %s has %d cars\n", u.firstName, u.lastName, (u.cars - u.carsSold))
}
// Pass by Reference: 
// *user means values of that object/struct can be changed
// using value with type string
func (u *user) SetFirstName(value string) {// Acts like setter in PHP or Java
	u.firstName = value
}

// Methods that return data
func (u user) GetFirstName() string{// Acts like getter in PHP or Java
	return u.firstName
}

