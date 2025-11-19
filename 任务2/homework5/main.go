package main

import "fmt"

/*
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
type Shape interface {
	Area() int
	Perimeter() int
}
type Rectangle struct {
	rec int
}

// Perimeter implements Shape.
func (r Rectangle) Perimeter() int {
	r.rec = r.rec + 150
	return r.rec
}

func (r Rectangle) Area() int {
	r.rec = r.rec + 100
	return r.rec
}

// Perimeter implements Shape.
// func (r Rectangle) Perimeter() int {
// 	panic("unimplemented")
// }

type Circle struct {
	cir int
}

// Area implements Shape.
func (c Circle) Area() int {
	c.cir = c.cir + 250
	return c.cir
}

func (c Circle) Perimeter() int {
	c.cir = c.cir + 200
	return c.cir
}

func main() {
	r := Rectangle{rec: 1}
	c := Circle{cir: 5}
	var s Shape = c
	var s1 Shape = r
	fmt.Println("area", s.Area())
	fmt.Println("Perimeter", s.Perimeter())
	fmt.Println("area", s1.Area())
	fmt.Println("Perimeter", s1.Perimeter())

}
