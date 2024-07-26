package main

import (
	"fmt"
	"os/user"
)

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
