package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var  (
	aa = false
	bb = "hi"
	cc = 1231231

)

func variableZeroValue()  {
	var a int
	var b string
	fmt.Printf("%d   %q \n",a,b)
}
func  variableInitialValue(){
	var a,b int  =3,4
	var s string = "hello a b c "
	fmt.Println(a,b,s)
}
func variabkeTypeDeduction()  {
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)

}
func variableShorter()  {
	a,b,c,s := 3,4,true,"shorter"
	fmt.Println(a,b,c,s)
}
//
func euler()  {
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println("%.3f\n",cmplx.Exp(1i*math.Pi) + 1 )

}
func triangle()  {
	a, b := 3,4
	c := calcTriangle(a,b)
	fmt.Println(c)
}
func calcTriangle(a,b int )int  {
	var c  int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
//
func consts()  {
	const (
		filename = "abc.txt"
		a,b = 3,4

	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename,c)
}

func enums()  {
	const(
		cpp =  iota
		javascrpit
		_
		python
		java
		golang
		php

	)
	const(
		aa = 1 << (10 * iota)
		bb
		cc
		dd
		ee
	)
	fmt.Println(cpp,javascrpit,python,java,golang,php)
	fmt.Println(aa,bb,cc,dd,ee)
	fmt.Println(1 << (10 *1 ))
}
func shifting (a int ) int {
	a = a << 1
	a = a >> 1
	return a
}
func main()  {
	variableZeroValue()
	variableInitialValue()
	variabkeTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,cc)

	euler()
	triangle()
	consts()
	enums()


}
/*
result
0   ""
3 4 hello a b c
3 4 true def
3 4 true shorter
false hi 1231231
(0+1.2246467991473515e-16i)
(0+1.2246467991473515e-16i)
%.3f
 (0+1.2246467991473515e-16i)
5
abc.txt 5
0 1 3 4 5 6
1 1024 1048576 1073741824 1099511627776
1024

 */
