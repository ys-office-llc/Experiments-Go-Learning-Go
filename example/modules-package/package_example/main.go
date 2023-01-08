package main

import (
	"fmt"

	"example.co.jp/package_example/aws/s3"          // パッケージ s3
	print "example.co.jp/package_example/formatter" // パッケージprint
	"example.co.jp/package_example/math"            // パッケージmath
)

func main() {
	num := math.Double(2)       // パッケージmath（math/math.go）
	output := print.Format(num) // パッケージprint（formatter/formatter.go）
	r := s3.PutObject("image.png")
	fmt.Println(output)
	fmt.Println(r)
}
