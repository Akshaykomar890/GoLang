package main

// func add(a int,b int) int{
// 	return a+b
// }

// func add(a ,b int) int{
// 	return a+b
// }

//Return multiple values
// func getStrings()(string,string,string,bool){
// 	return "Kotlin","Java","Go",true
// }

// func processIt(fn func(a int) int){
// 	fn(1)
// }

 func processIt() func(a int) int{
	return func (a int) int  {
		return 2
	}
 }


func main(){

	// fn:=func (a int) int  {
	// 	return 2
	// }
	fun :=processIt()
	fun(6)


	//fmt.Println(add(10,20))
	// lang1,lang2,lang4,lang5 := getStrings()
	//fmt.Println(getStrings())
	// fmt.Println(lang1,lang2,lang4,lang5)

	//If we don't use a any value we can use "_"
	// lang1,lang2,lang4,_ := getStrings()
	// fmt.Println(lang1,lang2,lang4)




}