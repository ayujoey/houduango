package main
import "fmt"

func main() {
	var a int
	var b int
	var ch string
	fmt.Println("输入你的算式")
	fmt.Scan(&a,&ch,&b)
	switch ch {
	case "+":
		fmt.Println(a, ch, b, "=", a+b)
	case "-":
		fmt.Println(a, ch, b, "=", a-b)
	case "*":
		fmt.Println(a, ch, b, "=", a*b)
	case "/":
		fmt.Println(a, ch, b, "=", a/b)
	}
}

