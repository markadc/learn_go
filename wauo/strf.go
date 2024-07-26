package wauo

import (
	"fmt"
	"strings"
)

// FormatString python的format
func FormatString(template string, values ...interface{}) string {
	for _, value := range values {
		template = strings.Replace(template, "{}", fmt.Sprintf("%v", value), 1)
	}
	return template
}

// PyPrint python的format输出
func PyPrint(template string, values ...interface{}) {
	result := FormatString(template, values...)
	fmt.Println(result)
}
