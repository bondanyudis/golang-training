package main

// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )

// func main() {
// 	// fmt.Println("")
// 	// rect := Rect{
// 	// 	width:  5.0,
// 	// 	height: 5.0,
// 	// }
// 	// cir := Circle{5.0}
// 	// fmt.Printf("Area rectangle = %0.1f\n", rect.Area())
// 	// fmt.Printf("Area circle = %0.1f\n", cir.Area())
// 	// fmt.Printf("Area circle = %0.1f\n", rect.area)
// 	// fmt.Println("")
// 	explain("Hello world")
// 	explain(52)
// 	explain(true)

// }
// func explain(i interface{}) {
// 	switch i.(type) {
// 	case string:
// 		fmt.Println("i stored string", strings.ToUpper(i.(string)))
// 	case int:
// 		fmt.Println("i stored integer", i)
// 	default:
// 		fmt.Println("i stored something else", i)
// 	}
// }

// type Rect struct {
// 	width  float64
// 	height float64
// 	area   float64
// }
// type Circle struct {
// 	radius float64
// }

// func (r Rect) Area() float64 {
// 	return r.width * r.width
// }
// func (r *Rect) Area1() {
// 	r.area = r.width * r.width
// }
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
