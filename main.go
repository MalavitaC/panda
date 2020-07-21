package main

import (
	"fmt"
	_ "panda/router"
	"reflect"

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

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal) GetSub(name string) {
	fmt.Printf("%v完成了减法运算 %v-%v=%v", name, cal.Num1, cal.Num2, cal.Num1-cal.Num2)
}

func main() {
	// app := gin.Default()
	// router.Register(app)
	// app.Run("0.0.0.0:9000")
	// fmt.Println("端口:", 9000)

	cal := Cal{}

	ct := reflect.TypeOf(cal)
	cv := reflect.ValueOf(&cal)
	for i := 0; i < cv.Elem().NumField(); i++ {
		fmt.Printf("结构体的字段%v: %v\n", i, ct.Field(i).Name)
	}
	cv.Elem().Field(0).SetInt(8)
	cv.Elem().Field(1).SetInt(3)

	var nameSilce []reflect.Value
	nameSilce = append(nameSilce, reflect.ValueOf("tom"))
	cv.MethodByName("GetSub").Call(nameSilce)

	// close(reschan)

	// for v := range reschan {
	// 	fmt.Println(v)
	// }

}
