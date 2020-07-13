package main

import (
	"fmt"
	_ "panda/router"

	_ "github.com/gin-gonic/gin"
)

func writechan(numchan chan int) {

	for i := 0; i < 2000; i++ {
		numchan <- i
	}
	close(numchan)
}

func readchan(numchan chan int, reschan chan int, exitchan chan bool) {
	for {
		num, ok := <-numchan
		if !ok {
			break
		}
		res := 0
		for i := 0; i < num; i++ {
			res += i
		}
		reschan <- res
	}
	exitchan <- true
	close(exitchan)
}

func main() {
	// app := gin.Default()
	// router.Register(app)
	// app.Run("0.0.0.0:9000")
	// fmt.Println("端口:", 9000)

	numchan := make(chan int, 2000)
	reschan := make(chan int, 2000)
	exitchan := make(chan bool, 1)

	go writechan(numchan)
	for i := 0; i < 8; i++ {
		go readchan(numchan, reschan, exitchan)
	}

	for {
		_, ok := <-exitchan
		if !ok {

			close(reschan)
			break
		}
	}

	for v := range reschan {
		fmt.Println(v)
	}

}
