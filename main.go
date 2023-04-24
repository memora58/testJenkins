package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	resp, err := http.Post("https://open.feishu.cn/open-apis/bot/v2/hook/f662967e-7ec7-4cf2-b438-13b94bd4f85c",
		"application/json",
		strings.NewReader("{\"msg_type\":\"text\",\"content\":{\"text\":\"构建成功\"}}"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}
