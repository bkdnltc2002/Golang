package main

import (
	"fmt"
	"os"
	// "os/signal"
	// "syscall"
	"bufio"
)



func main() { 
	// method 1
	// sigs := make(chan os.Signal, 1)

	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// done := make(chan bool,1)

	// go func ()  {
	// 	sig := <-sigs
	// 	fmt.Println()
	// 	fmt.Println(sig)
	// 	done <-true
	// }()
	// fmt.Println("awaiting signal")
	// <-done
	// fmt.Println("exiting")

	// method 2
	fmt.Println("Press Enter to stop")
	reader := bufio.NewReader(os.Stdin)

	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println("Program stopped")
}
