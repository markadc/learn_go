package main

import (
	. "demo/wauo"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Send(url string, headers map[string]string, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}

	// request
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return nil, e
	}

	// headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// response
	response, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	return response, nil
}

func main() {
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.0.0"
	headers := map[string]string{"User-Agent": ua}
	timeout := 3 * time.Second

	url := "https://www.baidu.com"
	response, e := Send(url, headers, timeout)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer response.Body.Close()

	body, e := io.ReadAll(response.Body)
	text := string(body)
	if e != nil {
		PyPrint("读取错误 | {}", e)
		return
	}
	fmt.Println(text)
	fmt.Println(len(text))
	fmt.Println(response.StatusCode)

}
