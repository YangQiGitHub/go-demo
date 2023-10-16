package fmtstd

import (
	"fmt"
	"os"
)

// DOC: https://pkg.go.dev/fmt
func FmtDemoFun() {
	// basicStdoutPrint()
	// printStruct()
	// printBoolean()
  // printInt()
  // printFloat()
  // printString()
  // printFloatFormat()
  // testFprint()
  // testSprint()
  // testErrorf()
  demo()
}

func basicStdoutPrint() {
	fmt.Print("Stdout print ", "without a newline.\n")
	fmt.Println("Stdout print,", "Spaces are always added between operands and a newline is appended. ")
	fmt.Printf("Stdout print without a newline. %v\n", 88)
}

type person struct {
	name   string
	age    int
	gender string
}

type student struct {
	id    int
	grade string
	*person
}

func printStruct() {
	xiaoMing := student{
		id:    1,
		grade: "No.1",
		person: &person{
			name: "LiMing",
		},
	}

	xiaoLiang := &student{
		id:    2,
		grade: "No.2",
		person: &person{
			name: "LiuLiang",
		},
	}

	fmt.Printf("%v\n", xiaoMing)
	fmt.Printf("%+v\n", xiaoMing)
	fmt.Printf("%#v\n", xiaoMing)
	fmt.Printf("%T\n", xiaoMing)
	fmt.Printf("%%\n")
	fmt.Println("--------------------")
	fmt.Printf("%v\n", xiaoLiang)
	fmt.Printf("%+v\n", xiaoLiang)
	fmt.Printf("%#v\n", xiaoLiang)
	fmt.Printf("%T\n", xiaoLiang)
	fmt.Printf("%%\n")
	fmt.Println("--------------------")
	fmt.Printf("%v\n", xiaoMing.person)
	fmt.Printf("%+v\n", xiaoMing.person)
	fmt.Printf("%#v\n", xiaoMing.person)
}

func printBoolean() {
	fmt.Printf("%t\n", true)
	fmt.Printf("%t\n", false)
}

func printInt() {
  num := 180
  fmt.Printf("%b\n", num)
  fmt.Printf("%c\n", num)
  fmt.Printf("%d\n", num)
  fmt.Printf("%o\n", num)
  fmt.Printf("%O\n", num)
  fmt.Printf("%q\n", num)
  fmt.Printf("%x\n", num)
  fmt.Printf("%X\n", num)
  fmt.Printf("%U\n", num)
}

func printFloat() {
  floatNum := 18.54
  fmt.Printf("%b\n", floatNum)
  fmt.Printf("%e\n", floatNum)
  fmt.Printf("%E\n", floatNum)
  fmt.Printf("%f\n", floatNum)
  fmt.Printf("%F\n", floatNum)
  fmt.Printf("%g\n", floatNum)
  fmt.Printf("%G\n", floatNum)
  fmt.Printf("%x\n", floatNum)
  fmt.Printf("%X\n", floatNum)
}

func printString() {
  str := "我是字符串"
  fmt.Printf("%s\n", str)
  fmt.Printf("%q\n", str)
  fmt.Printf("%x\n", str)
  fmt.Printf("%X\n", str)
  fmt.Println("----------------")
  b := []byte{65, 66, 67}
  fmt.Printf("%s\n", b)
  fmt.Printf("%q\n", b)
  fmt.Printf("%x\n", b)
  fmt.Printf("%X\n", b)
}

func printFloatFormat() {
  num := 13.14
  fmt.Printf("%f\n", num)
  fmt.Printf("%10f\n", num)
  fmt.Printf("%.2f\n", num)
  fmt.Printf("%10.2f\n", num)
  fmt.Printf("%10.f\n", num)
  fmt.Printf("%f\n", num)
}

func testFprint() {
  fmt.Fprint(os.Stdout, "向标准输出打印\n")
  fmt.Fprintln(os.Stdout, "向标准输出中打印")
  fmt.Fprintf(os.Stdout, "%s 向标准输出打印\n", "哈哈")
}

func testSprint() {
  host := "localhost"
  port := 6379
  addr := fmt.Sprintf("%s:%d", host, port)
  fmt.Println(addr)
}

func testErrorf() {
  err := fmt.Errorf("用户名格式不正确: %s", "#赛帆")
  if err != nil {
    panic(err)
  }
}

func demo() {
  fmt.Fprintf(os.Stdout, "bb %s %s", "cc %s", "dd")
}
