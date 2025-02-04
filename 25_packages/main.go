package main

import (
	"fmt"
	"os/user"

	"github.com/Akshaykomar890/GolangProject/auth"
)




func main(){

	auth.LoginWithCred("akshay","pass")
	session:=auth.GetSession()
	fmt.Println(session)

	user:=user.User{
		Name: "akshay",
	}
	fmt.Println(user.Name)


}