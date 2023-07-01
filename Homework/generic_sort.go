package main

import(
	"fmt"
)

type Form interface{
	int | uint32| uint64
}

func sort_slice[F Form](s []F) []F  {

	n:= len(s)
	for i:=0;i<n-1;i++{
		for j:=0;j<n-i-1;j++{
			if s[j]>s[j+1]{
				s[j],s[j+1]=s[j+1],s[j]
			}
		}
	}
	return s
}

func main(){
	int_slice :=[]int{5,2,7,1,9,19,34,23,54,24,64,67,53}
	uint32_slice := []uint32{6,0,1,5,9,4,2,3,7,8}
	uint64_slice := []uint64{9,8,5,0,3,4,6,7,2,1}
	fmt.Println("Before sorting: ")
	fmt.Println("int: ",int_slice)
	fmt.Println("uint32: ",uint32_slice)
	fmt.Println("uint64: ",uint64_slice)
	fmt.Print()
	fmt.Println("After sorting: ")
	fmt.Println("int: ",sort_slice(int_slice))
	fmt.Println("uint32: ",sort_slice(uint32_slice))
	fmt.Println("uint64: ",sort_slice(uint64_slice))
}