package main
import "fmt"
func main() {
	var a int = 42
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
	var a int = 42
	var b *int = &a
	fmt.Println(a, b, *b)
	a = 27
	fmt.Println(a, *b)
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1] - 4
	fmt.Println("%v %p %p\n", a, b, c)
	var ms *myStruct
	ms = new(myStruct)
	ms = myStruct{foo: 42}
	(*ms).foo = 42
	// compiling helping us out
	ms.foo = 42
	fmt.Println(ms, ms.foo, (*ms).foo)
}
type myStruct struct {
	foo int
}
