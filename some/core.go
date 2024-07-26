package some

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReqUrl(url string) (string, error) {
	// 创建一个 HTTP 客户端
	client := http.Client{}
	// 发送 GET 请求
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET ERROR => %v\n", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("READ ERROR => %v\n", err)
	}

	// 返回响应内容
	return string(body), nil
}
