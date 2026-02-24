package example

import "fmt"

// SayHello 打印一行字符串
func SayHello() {
	fmt.Println("Hello World")
}

// SayGoodbye 打印两行字符串
func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

// PrintNames 打印学生姓名
func PrintNames() {
	students := make(map[int]string, 4)
	students[1] = "Jim"
	students[2] = "Bob"
	students[3] = "Tom"
	students[4] = "Sue"
	for _, value := range students {
		fmt.Println(value)
	}
}

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
		return
	}
	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}

// 检测单行输出
func ExampleSayHello() {
	SayHello()
	// OutPut: Hello World
}

// 检测多行输出
func ExampleSayGoodbye() {
	SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	PrintNames()
	// Unordered output:
	// Jim
	// Bob
	// Tom
	// Sue
}

func ExampleGreetingWithoutParameter() {
	Greeting("nobody")
	// OutPut:
	// Nobody to say hi.
}

func ExampleGreetingWithParameter() {
	Greeting("hello:", "Joe", "Anna", "Eileen")
	// OutPut:
	// hello: Joe
	// hello: Anna
	// hello: Eileen
}
