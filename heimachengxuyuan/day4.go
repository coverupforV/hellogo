package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// 这是给struct Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}

// 定义一个Point切片类型 Path
type Path []Point

// 方法的接收器 是Path类型数据, 方法的选择器是TranslateBy(Point, bool)
func (path Path) TranslateBy(another Point, add bool) {
	var op func(p, q Point) Point //定义一个 op变量 类型是方法表达式 能够接收Add,和 Sub方法
	if add == true {
		op = Point.Add //给op变量赋值为Add方法
	} else {
		op = Point.Sub //给op变量赋值为Sub方法
	}

	for i := range path {
		//调用 path[i].Add(another) 或者 path[i].Sub(another)
		path[i] = op(path[i], another)
		path[i].Print()
	}
}

func main() {

	points := Path{
		{10, 10},
		{11, 11},
	}

	anotherPoint := Point{5, 5}

	points.TranslateBy(anotherPoint, false)

	fmt.Println("------------------")

	points.TranslateBy(anotherPoint, true)
}
