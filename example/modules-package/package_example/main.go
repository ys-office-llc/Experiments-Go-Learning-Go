package main

import (
	"fmt"

	"example.co.jp/package_example/aws/s3"
	print "example.co.jp/package_example/formatter"         // パッケージprint
	print_sub "example.co.jp/package_example/formatter/sub" // パッケージ print_sub
	"example.co.jp/package_example/math"                    // パッケージmath
)

func main() {
	num := math.Double(2)       // パッケージmath（math/math.go）
	output := print.Format(num) // パッケージprint（formatter/formatter.go）
	output2 := print.Print(num * 2)
	output3 := print_sub.PrintSub(num * 3)
	o := s3.PutObject("image.png")
	fmt.Println(output)
	fmt.Println(output2)
	fmt.Println(output3)
	fmt.Println(o)
}
