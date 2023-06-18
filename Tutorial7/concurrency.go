package main 

import(
	"fmt"
	"time"
	"golang.org/x/tour/tree"
)

func say(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int,c chan int){
	sum :=0
	for _, v:= range s{
		sum+=v
	}
	c <- sum
}

func fibonacci(c, quit chan int){
	x,y :=0,1
	for {
		select {
		case c<-x:
			x,y=y,x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int){
	if t.Left!= nil{
		Walk(t.Left,ch)
	}
	ch <-t.Value
	if t.Right!=nil{
		Walk(t.Right,ch)
	}
}

func Same(t1,t2 *tree.Tree)bool{
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1,ch1)
	go Walk(t2,ch2)

	for i:=0;i<10;i++{
		if<-ch1!=<-ch2{
			return false
		}
	}
	return true
}

func main (){
	// go say("world")
	// say("hello")
	// s :=[]int{7,2,8,-9,4,0}

	// c :=make(chan int)
	// go sum(s[:len(s)/2],c)
	// go sum(s[len(s)/2:],c)
	// x,y :=<-c,<-c

	// fmt.Println(x,y,x+y)

	// ch := make(chan int,2)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// c:= make(chan int)
	// quit := make(chan int)
	// go func ()  {
	// 	for i:=0;i<10;i++{
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci(c,quit)

	// tick := time.Tick(100 *time.Millisecond)
	// boom := time.After(500* time.Millisecond)
	// for {
	// 	select {
	// 	case<-tick:
	// 		fmt.Println("tick.")
	// 	case<-boom:
	// 		fmt.Println("BOOM!")
	// 		return
	// 	default:
	// 		fmt.Println("    .")
	// 		time.Sleep(50*time.Millisecond)
	// 	}
	// }

	ch := make(chan int)
	go Walk(tree.New(1),ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Equivalent Binary Trees?", Same(tree.New(1),tree.New(1)))
	fmt.Println("Equivalent Binary Trees?", Same(tree.New(1),tree.New(2)))

}