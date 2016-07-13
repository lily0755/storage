package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "中国 中国"
	tt,_ := urlEncoded(str)
	fmt.Println("编码后的字符",tt)
}


func urlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
