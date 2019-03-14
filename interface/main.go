package main

import (
	"math"
	"fmt"
	"github.com/JimmyHongjichuan/grpc_test/interface/compose"
)

// 几何图形的接口(interface)
type geometry interface {
	area(in float64) float64
	perim(in float64) float64
}

// rect 和 circle 实现 geometry 接口
type squr struct {

}
type circle struct {
	radius float64
}


// 实现了接口中声明的所有方法即实现了该接口，这里 rects 实现了 geometry 接口
func (r squr) area(in float64) float64 {
	return  in*in
}
func (r squr) perim(in float64) float64 {
	return 4*in
}

// circles 实现 geometry 接口
func (c circle) area(in float64) float64 {
	return math.Pi * in * in
}
func (c circle) perim(in float64) float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry, in float64) {
	fmt.Println(g)
	fmt.Println(g.area(in))
	fmt.Println(g.perim(in))
}
type Node struct{
	compose.BaseService
}


func (node Node)GetMsg() string {
	return node.BaseService.Msg
}

func main() {
	node := &Node{}
	node.SetData("hello")
	fmt.Println(node.GetMsg())
	r := squr{}
	c := circle{radius: 5}
	// The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	// rect 和 circle 结构类型都实现了 geometry 接口，所以我们能够把它们的对象作为参数传给 measure 函数
	// 在 measure 函数内部，参数对象调用了所属类型实现的接口方法。
	measure(r, 4)
	measure(c, 4)
}