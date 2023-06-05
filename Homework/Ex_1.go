package main

import(
	"fmt"
)

func Sqrt(i []int) []int{
	result := make([]int,len(i))
	for index,num := range i {
		result[index] = num*num
	}
	return result
}

func Update_Sqrt(i []int) []int{
	for index,num := range i {
		i[index] = num*num
	}
	return i
}

func Bubble_Sort(arr []int){
	fmt.Println("Before bubble sort: ",arr)
	n:= len(arr)
	count :=0
	for i:=0;i<n-1;i++{
		for j:=0;j<n-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
				count +=1
			}
		}
	}
	fmt.Println("After bubble sort: ",arr)
	fmt.Println("Times of comparing: ",count)
}

func main(){
	input := []int{4,6,7,10,3}
	//Print all element ^2
	fmt.Println("----- Print all elements ^2-----")
	fmt.Println(Sqrt(input))
	fmt.Println(input)
	fmt.Println("----- Modify all elements to ^2-----")
	fmt.Println(Update_Sqrt(input))
	fmt.Println(input)
	fmt.Println("----- Bubble Sort-----")
	Bubble_Sort(input)
}