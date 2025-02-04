package main

import (
	"fmt"

	"github.com/Akshaykomar890/GolangProject/auth"
)




func main(){

	auth.LoginWithCred("akshay","pass")
	session:=auth.GetSession()
	fmt.Println(session)

	// user:=user.User{
	// 	Name: "akshay",
	// }
	// fmt.Println(user.Name)

	//color.Green(user.Name)


}