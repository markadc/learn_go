package main

import (
	"fmt"
	"os/user"
)

// 输出当前电脑的用户信息
func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(u.Uid)
	fmt.Println(u.Gid)
	fmt.Println(u.Username)
	fmt.Println(u.HomeDir)
	fmt.Println(u.Name)

}
