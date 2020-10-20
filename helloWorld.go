package main

import (
	"fmt"
)

func main() {

	var a int = 21
	var b int = 10
	var c int

	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
	a = 60
	c = a << 6 // 128 64  32 16 8 4 2 1     0011 1100     1111 0000 0000
	fmt.Printf("第八行 -c 的值为 ： %d \n", c)
	fmt.Printf("sum (a + b) = % d\n", sum(a, b))

	/*
	   第七行 - a 的值为 20
	   第八行 -c 的值为 ： 3840
	   a = 60
	   b = 10
	    age[3] =54.360001
	*/

	//指针   还有空指针
	var ip *int
	ip = &a
	fmt.Printf("ip :  %x \n", ip)

	//结构体的使用
	type popleObj struct {
		name string
		age  int
		sex  string
	}

	fmt.Println(popleObj{"xiao wang", 23, "man"})
	fmt.Println(popleObj{name: "xiao li ", age: 34})
	/*
	   数组中的元素：  1.300000
	   数组中的元素：  2.000000
	   数组中的元素：  5.025000
	   数组中的元素：  54.360001
	   数组中的元素：  0.000000
	   sum (a + b) =  70
	   ip :  c0000120a0
	   {xiao wang 23 man}
	   {xiao li  34 }
	*/

	go test_func(9)
	fmt.Println("12312")

}

//  函数
func sum(a, b int) int {
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n ", b)

	var age = [5]float32{1.3, 2.0, 5.025, 54.36} //数组
	fmt.Printf("age[3] =%f\n", age[3])
	var i int
	for i = 0; i < 5; i++ {
		fmt.Printf("数组中的元素：  %f \n", age[i])

	}

	return a + b
}
