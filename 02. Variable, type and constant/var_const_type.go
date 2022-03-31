package main

import "fmt"

//globalvar := 2000032  not allowed

const LoginToken string = "ghhhrajh"

func main() {
	var username string = "Ravindra"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("variable is of type: %T \n", isloggedin)

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("variable is of type: %T \n", smallvalue)

	var anothervariableint int
	fmt.Println(anothervariableint)
	fmt.Printf("variable is of type: %T \n", anothervariableint)

	var anothervariablestring string
	fmt.Println(anothervariablestring)
	fmt.Printf("variable is of type: %T \n", anothervariablestring)

	var x = 2
	fmt.Println(x)
	fmt.Printf("variable is of type: %T \n", x)

	y := 30.000078
	fmt.Println(y)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
