package main

//Slices are Dynamic Array
func main(){

	//unintiallized slice is null
	//var nums[]int
	//fmt.Println(nums == nil)
	//fmt.Println(len(nums))
	
	//Another way
	//var nums = make([]int,0,5) //5 -> intial capacity
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	//Another way to make Slice
	//nums:=[]int{}


	// nums = append(nums, 1)
	// nums[0] = 3
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))


	//Copy
	// var nums = make([]int,0,5)
	// nums = append(nums, 2)
	// var nums2 = make([]int,len(nums))
	// //copy fun
	// copy(nums2,nums)

	// fmt.Println(nums,nums2)

	// var nums = []int{1,2,3}
	// //fmt.Println(nums[0:2]) //1,2
	// //or
	// //fmt.Println(nums[:2]) //1,2
	// fmt.Println(nums[0:])  //123 or 1: -> 2,3

	//Compare
	// var nums1 = []int{1,2}
	// var nums2 = []int{1,2}

	// fmt.Println(slices.Equal(nums1,nums2))






}