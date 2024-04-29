package main
import "fmt"
func main() {
	var n int
	var a,b string = " ",""
	fmt.Scan(&n)
	for n>=1{
		b = ""
		fmt.Print(a)
		a = "*" + a
		for x:=n ; x>1 ; x--{
			b += "*"
		}
		fmt.Println(b)
		n--
	}
}