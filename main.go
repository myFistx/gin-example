package main

import (
	"fmt"
	"gin-blog/pkg/setting"
)

func main()  {
	setting.Init()
	fmt.Println(setting.HttpPort)
	fmt.Println(setting.JwtSecret)
}