// https://trap.jp/post/1445/
package main

import "fmt"

// Circle（円）という構造体を定義 *1
type Circle struct {
	Radius int
}

// Circleのメソッドを定義 *2
func (c Circle) GetArea() int {
	return 3 * c.Radius * c.Radius
}

// Square（正方形）という(ry
type Square struct {
	Height int
}

func (s Square) GetArea() int {
	return s.Height * s.Height
}

// Figure（図形）インターフェイスを定義 *3
type Figure interface {
	GetArea() int
}

// Figureの面積を表示する関数を定義 *4
func DisplayArea(f Figure) {
	fmt.Printf("面積は%vです\n", f.GetArea())
}

// 実行される部分 *5
func main() {
	circle := Circle{Radius: 1}
	DisplayArea(circle)

	square := Square{Height: 3}
	DisplayArea(square)
}
