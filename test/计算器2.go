package main

import (
	"bufio"
	"fmt"
	"github.com/Knetic/govaluate"
	"os"
	"strings"
)

func start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入数学表达式（例如 1+2/2*2）：")
		expressionStr, _ := reader.ReadString('\n')
		fmt.Println("你输入=", expressionStr)
		expressionStr = strings.TrimSpace(expressionStr)
		if expressionStr == "" {
			fmt.Println("请输入有效的数学表达式。")
			continue
		}
		expression, err := govaluate.NewEvaluableExpression(expressionStr)
		fmt.Println("解析表达式后=", expression)
		if err != nil {
			fmt.Println("输入错误，请输入有效的数学表达式。")
			continue
		}
		fmt.Println("准备计算了...")

		// 计算结果
		result, err := expression.Evaluate(nil)
		if err != nil {
			fmt.Println("计算错误：", err)
			continue
		}

		fmt.Printf("结果是：%v\n", result)

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

func main() {
	start()
}
