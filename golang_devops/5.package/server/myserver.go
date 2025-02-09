package server

import "fmt"

type MyServer struct {
	Address string
	Config  string /// wszystko z dużej aby było eksportowane
}

func init() {
	fmt.Println("wywolana init()")
}