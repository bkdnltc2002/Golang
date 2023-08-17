package main

// "os/signal"
// "syscall"

func main() {
	// method 3 (using for loop)
	// go func() {
	// 	count := 0
	// 	for {
	// 		count += 1
	// 		fmt.Println(count)
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }()
	// for {
	// 	time.Sleep(time.Second * 1)
	// }

	// method 4 (loop in goroutine)
	// go func() {
	// 	count := 0
	// 	for {
	// 		count += 1
	// 		fmt.Println(count)
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }()
	// ch := make(chan bool)
	// <-ch

	//web crawler (url_title_content)
	
}
