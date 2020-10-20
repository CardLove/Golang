package main


import(
	"fmt"
	"time"
)

func main() {


	for i := 0; i < 100; i++ {
		test_goroute(i)
	}

	time.Sleep(time.Second)
	fmt.Println("main"+ "finsh")
	var str string  = "wlh"

	var numa , numb  int  =   2, 4
	fmt.Println(str + "numa:"+ numa + "numb:" +numb)
}

func test_goroute(a int) {

	fmt.Println(a)
	fmt.Println("sdfsdfs")
}