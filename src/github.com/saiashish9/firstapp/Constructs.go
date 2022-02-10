package main

import "fmt"

func main() {

	if !false {
		fmt.Println("Test")
	}
	s := map[string]int{
		"a": 1,
		"b": 2,
	}
	if c, d := s["a"]; d {
		fmt.Println(c)
	}
	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("same")
	} else if 0 == 0 {
		fmt.Println("diff")
	}
	switch i := 2 + 3; i {
	case 1, 2:
		fmt.Println("one")
	case 3:
		fmt.Println("two")
	default:
		fmt.Println("neither one nor two")
	}
	i := 2
	switch {
	case i <= 10:
		fmt.Println("<=0")
		fallthrough
	case i >= 20:
		fmt.Println(">=0")
	}
	var i interface{} = 1
	i = 2.2
	i = true
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case float32:
		fmt.Println("float32")
		fmt.Println(i)
	}
	for i := 0; i < 5; i += 2 {
		fmt.Println(i)
	}
	// i++ X
	// i=0 , j=0 X
	for i, j := 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Label

	// Loop:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Println(i * j)
	// 			if i*j >= 3 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	s := []int{1, 2, 3}
	fmt.Println(s)

	for k, v := range s {
		fmt.Println(k, v)
	}

	x := map[string]int{
		"a": 1,
	}

	for _, p := range x {
		fmt.Println(p)
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

}
