package main
import "fmt"
type Doctor struct {
	number     int
	actorName  string
	companions []string
}
type A struct {
	name   string
	origin string
}
type B struct {
	A
	Canfly bool
}
func main() {
	s := map[string]int{
		"a": 1,
		"b": 2,
	}
	s["a"] = 3
	delete(s, "a")
	_, b := s["a"]
	// no need to make use of _ again
	fmt.Println(s, b)
	a := Doctor{
		number:    3,
		actorName: "Sai",
		companions: []string{
			"a",
			"b",
		},
	}
	b := Doctor{
		3,
		"Sai",
		[]string{
			"a",
			"b",
		},
	}
	a := struct{ name string }{name: "Sai"}
	b := &a
	b.name = "Ashish"
	fmt.Println(a, b)
	// go doesn't supports traditional oop features
	// go doesn't have inheritance
	// fmt.Println(b.actorName)
	b := B{}
	b.Canfly = true
	b.name = "Sai"
	fmt.Println(b, b.name)
	type A struct {
		name   string
		origin string
	}
	type B struct {
		A
		destination string
	}
	b := B{
		A: A{"sai", "123"},
		destination: "delhi",
	}
	fmt.Println(b)
}
