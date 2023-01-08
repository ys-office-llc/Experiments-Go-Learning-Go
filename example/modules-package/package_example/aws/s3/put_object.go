package s3

import "fmt"

func PutObject(file string) string {
	return fmt.Sprintf("パスは %s です。", file)
}
