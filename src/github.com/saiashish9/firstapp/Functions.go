package main
import "fmt"
func main() {
	sayMsg("Go")
	greeting := "Hello"
	name := "Sai"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
	sum(1, 2, 3, 4)
	g := greeter{
		"hi", "Go",
	}
	g.greet()
	fmt.Println(g.name)
}
type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g)
	g.name = ""
}
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(result)
}
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Sai9"
	fmt.Println(*name)
}
func sayMsg(msg string) {
	fmt.Println(msg)
}