package main

import(
	"fmt"
)


func removeChar()string{
	word:="Hello World"
	var result string
	for _,letter := range word{
		if string(letter)!=" "{
			result = result + string(letter)
		}
	}
	return result
}

func removeChar1()string{
	word:="Hello World"
	return word[:5]+word[6:]
}

func removeChar2()string{
	word:="Xuân Thành"
	// runeslice := []rune(word)
	var result string
	for _,letter := range word{
		if string(letter)!=" "{
			result = result + string(letter)
		}
	}
	
	return result 
}

func removeChar3()string{
	word:="Xuân Thành"
	runeslice := []rune(word)
	runeslice = append(runeslice[:4],runeslice[5:]...)
	return string(runeslice)
}

func main(){
	// binary vs text ???
	// restfulAPI data transfer using text instead of binary, pros/cons
	// slices package using ???
	fmt.Println(removeChar())
	fmt.Println(removeChar1())
	fmt.Println(removeChar2())
	fmt.Println(removeChar3())


}