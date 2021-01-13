// package firstapp

//  i:=42 X
// var i int = 42
// local package access
// I
// global package access

// func main() {
// var i int
// i = 42
// var i int = 42
// i := 42
// var j float32 = 27
// go compiler can
// automatically figure out
// the datatype itself
// fmt.Println(i)
// fmt.Printf("%v ,%T", j, j)
// var a float32 = 42.5
// var j int
// j = int(a)
// we have to do explicit conversion ourself
// println(j)
// var p int = 32
// fmt.Println(p)
// var q string
// q = string(p)
// it will return a byte
// fmt.Printf("%v, %T", q, q)
// q = strconv.Itoa(p)
// fmt.Printf("%v, %T", q, q)
// var (
// 	x string = "A"
// 	y string = "B"
// )
// variable blocks
// var (
// 	i = 32
// )
// fmt.Println(x, y)
// fmt.Printf("%v ,%T", i, i)
// var n bool = true
// fmt.Printf("%v ,%T", n, n)
// m := 1 == 1
// // n := 2 == 2
// var n bool
// // default 0 value false
// fmt.Printf("%v , %T\n", n, n)
// fmt.Println(m, n)

// int8 +128 - 127
// int16 -32768 - 32767
// int32 int64
// var n uint16 = 42
// uint32 uint64 X
// uint8 -> byte common
//    a:=3 b:=6 c:=a%b
// fmt.Println(3 / 6)
// var i int8 = 6
// fmt.Println(3 + i)
// correct int(3) + i wrong
// int(3) + int(i) correct

// var i complex64 = 1 + 2i
// i = complex(5, 6)
// fmt.Println(i)
// fmt.Println(real(i), imag(i))

// s := "abcd"
// println("abcd"[2], s[2])
//  s[2]="j" X
// & | ^ &^ xnor << >>
//    s+s correct
// b := []byte(s)
// println(b)
// fmt.Printf("%v ", b)

// rune '' int32
// r := 'a'
// fmt.Println(r)
// var r rune = 'i'
// fmt.Println(r)

// constants

// const myConst int = 42
// fmt.Println(myConst)
//   compile time

// const c float32 = math.Sin(1.57) X
// fmt.Println(c)

// const a int16 = 27
// func main(){
// const a int = 14
// var b int = 27
// fmt.Printf("%v %T\n",a,a)  }

// iota counter we can use while creating a numeric
// constant

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// const (
// 	a2 = iota
// 	b2
// )

// var c2 int = a2
// fmt.Println(a2 == c2)

// const (
// 	_ = 5
// 	a
// 	b
// 	c
// )
// println(a, b, c)
// 5 5 5
// const (
// 	_ = iota + 5
// 	a
// 	b
// 	c
// )
// println(a, b, c)
// 6 7 8

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )
// fileSize := 4000000000.00
// fmt.Printf("%.2fGB", fileSize/GB)

// const (
// 	a = 1 << iota
// 	b
// 	c
// 	e
// )
// 1 2 4 16

// var d byte = a | b | c
// println(d)
// fmt.Printf("%b\n", d)
// fmt.Printf("%v\n", a&d == a)
// fmt.Printf("%v\n", a&e == a)

// grades := [3]int{97, 85}
// g := [...]int{1, 2}
// var students [3]string
// fmt.Println(students)
// students[0], students[1] = "a", "b"
// fmt.Println(students, len(students))
// fmt.Printf("%v \n%v \n", grades, g)

// var a [3][3]int = [3][3]int{
// 	[3]int{1, 0, 1},
// 	[3]int{0, 0, 1},
// 	[3]int{1, 0, 0}}
// fmt.Println(a)

// a := [...]int{1, 2, 3}
// b := a
// c := &a
// c[1] = 5
// b[1] = 5
// fmt.Println(a, b, c)

// slice

// a := []int{1, 2, 3}
// fmt.Println(a, a[:2])

// a := make([]int, 3, 1000)
// fmt.Println(len(a), cap(a))
// a = append(a, 1, 2)
// b := []int{}
// b = append(b, []int{2, 3, 4}...)
// fmt.Println(b)
// fmt.Println(a)

// a := []int{1, 2, 3, 4, 5}
// fmt.Println(a)
// b := a[:len(a)]
// fmt.Println(b)
// }

// _ = iota
// write only value
// we can declare variables twice in same scope
