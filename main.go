package main

import (
	"fmt"
	_ "panda/router"
	"time"

	_ "github.com/gin-gonic/gin"
)

func writechan(numchan chan int) {

	for i := 1; i <= 800000; i++ {
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
		for i := 0; i <= num; i++ {
			res += i
		}
		reschan <- res
	}
	exitchan <- true
}

func main() {
	// app := gin.Default()
	// router.Register(app)
	// app.Run("0.0.0.0:9000")
	// fmt.Println("端口:", 9000)

	numchan := make(chan int, 800000)
	reschan := make(chan int, 800000)
	exitchan := make(chan bool, 8)

	go writechan(numchan)

	start := time.Now()
	for i := 0; i < 8; i++ {
		go readchan(numchan, reschan, exitchan)
	}

	for i := 0; i < 8; i++ {
		<-exitchan
	}
	cost := time.Since(start)
	fmt.Printf("耗时[%v]", cost)

	// close(reschan)

	// for v := range reschan {
	// 	fmt.Println(v)
	// }

}
