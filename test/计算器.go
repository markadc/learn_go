package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("请输入第一个数字：")
		num1Str, _ := reader.ReadString('\n')
		num1, err := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)
		if err != nil {
			fmt.Println("输入错误，请输入有效的数字。")
			continue
		}

		fmt.Println("请输入运算符 (1: +, 2: -, 3: *, 4: /)：")
		operator, _ := reader.ReadString('\n')
		operator = strings.TrimSpace(operator)

		fmt.Println("请输入第二个数字：")
		num2Str, _ := reader.ReadString('\n')
		num2, err := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)
		if err != nil {
			fmt.Println("输入错误，请输入有效的数字。")
			continue
		}

		var result float64
		switch operator {
		case "1":
			result = num1 + num2
		case "2":
			result = num1 - num2
		case "3":
			result = num1 * num2
		case "4":
			if num2 == 0 {
				fmt.Println("除数不能为零。")
				continue
			}
			result = num1 / num2
		default:
			fmt.Println("无效的运算符，请输入 1, 2, 3, 或 4。")
			continue
		}

		fmt.Printf("结果是：%f\n", result)

		for {
			fmt.Println("是否继续计算？ (yes/no)")
			continueStr, _ := reader.ReadString('\n')
			continueStr = strings.TrimSpace(strings.ToLower(continueStr))
			if continueStr == "yes" {
				break
			} else if continueStr == "no" {
				fmt.Println("退出程序。")
				return
			} else {
				fmt.Println("无效输入，请输入 'yes' 或 'no'。")
			}
		}
	}
}
