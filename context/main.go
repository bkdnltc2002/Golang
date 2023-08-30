package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	go func (ctx context.Context){
		for {
			select{
			case<-ctx.Done():
				// if err:= ctx.Err(); err != nil {
				// 	fmt.Println("Get err")
				// }
				fmt.Println("Finished")
				close(ch)
				return
			default:
				fmt.Println("Hello")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	<-ch
}