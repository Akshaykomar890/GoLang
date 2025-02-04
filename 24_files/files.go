package main

import (
	"os"
)

func main(){
	// Access the file
	// f,error:=os.Open("example.txt")

	// if error != nil {
	// 	panic(error)
	// }

	// info,error:=f.Stat()

	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Println("file name:",info.Name())
	// fmt.Println("file or folder:",info.IsDir()) //Check folder
	// fmt.Println("file size:",info.Size())


	//Read File 1st method difficult
	// f,error:= os.Open("example.txt") //locate
	// if error != nil { //find error
	// 	panic(error)
	// }
	// defer f.Close()  //for closing file
	// buffer := make([]byte,100) //store data 
	// d,error:=f.Read(buffer) //first read and store data in buffer
	// if error != nil{
	// 	panic(error)
	// }
	// println("data is:",string(buffer),"lengith",d) 


	//2nd method easy method
	// data,error:=os.ReadFile("example.txt") //loads entire file in memory not best
	// if error != nil {
	// 	panic(error)
	// }
	// fmt.Println(string(data)) //it prints byte so convert it



	//Read Folders
	//dir,error:=os.Open(".")//current folder
	// dir,error:=os.Open("../")//all folders

	// if error != nil {
	// 	panic(error)
	// }
	// defer dir.Close()
	// fileinfo,error:=dir.ReadDir(100) //returns slice

	// for _,fi:=range fileinfo{
	// 	fmt.Println(fi.Name(),fi.IsDir()) //check file or folder

	// }



	//File Write 
	//Create File
	file,error:=os.Create("example2.txt")
	if error != nil{
		panic(error)
	}
	defer file.Close()
	//write 1st method
	// file.WriteString("Hi Golang")
	// file.WriteString("Nice lang") //it appends

	//2nd method
	// byte := []byte("hello go lang")
	// file.Write(byte)

	
	






















}